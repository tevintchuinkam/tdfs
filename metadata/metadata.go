package metadata

import (
	context "context"
	"fmt"
	"io/fs"
	"log/slog"
	"math/rand"
	"path"
	"strings"
	"sync"
	"time"

	"log"
	"net"

	"github.com/google/uuid"
	"github.com/tevintchuinkam/tdfs/chunks"
	"google.golang.org/grpc"
)

func New(port int) *MetaDataServer {
	return &MetaDataServer{
		port: port,
	}
}

func (s *MetaDataServer) Start() {
	// accept connections
	addr := fmt.Sprintf(":%d", s.port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", addr, err)
	}
	grpcServer := grpc.NewServer()
	RegisterMetadataServiceServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func (s *MetaDataServer) RegisterChunkServer(port int) error {
	// ping the server
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("could not connect. err: %v", err)
	}
	defer conn.Close()
	c := chunks.NewChunkServiceClient(conn)
	// ping the server
	challenge := rand.Int63()
	resp, err := c.Ping(context.Background(), &chunks.PingRequest{Challenge: challenge})
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	if challenge != resp.Challenge {
		return fmt.Errorf("chunk server failed challenge: %d !=  %d (expected)", resp.Challenge, challenge)
	}
	s.muChunk.Lock()
	s.chunkServers = append(s.chunkServers, &c)
	s.muChunk.Unlock()
	return nil
}

func (s *MetaDataServer) GetMetadata(ctx context.Context, in *MetadataRequest) (*MetadataResponse, error) {
	slog.Info("received chunk data request", "filename", in.Filename, "chunk_index", in.ChunkIndex)
	return &MetadataResponse{
		ChunkHandle: "abcde",
		Url:         "chunk:9000",
	}, nil
}

type directory struct {
	name    string
	prev    *directory
	subDirs []*directory
	files   map[string]*File
}

func (d *directory) String() string {
	if d == nil {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(d.name)
	for _, file := range d.files {
		sb.WriteString(fmt.Sprintf("\n\t├── %s", file.name))
	}
	return sb.String()
}

type chunkID string

type Chunck struct {
	ID   chunkID
	Mu   sync.Mutex
	data []byte
}

type File struct {
	name   string
	chunks []*Chunck
}

type FileInfo struct {
}

func (f *FileInfo) Name() string {
	return "file"
}
func (f *FileInfo) Size() int64 {
	return 0
}
func (f *FileInfo) Mode() fs.FileMode {
	return 0
}
func (f *FileInfo) ModTime() time.Time {
	return time.Now()
}
func (f *FileInfo) IsDir() bool {
	return false
}

func (f *FileInfo) Sys() any {
	return struct{}{}
}

// File implements fs.File
func (f *File) Close() error {
	return nil
}
func (f *File) Stat() (fs.FileInfo, error) {
	return &FileInfo{}, nil
}

func (f *File) Read(b []byte) (int, error) {
	return 0, nil
}

type MetaDataServer struct {
	UnimplementedMetadataServiceServer
	port         int
	muChunk      sync.Mutex
	chunkServers []*chunks.ChunkServiceClient
	fileMap      map[string][]chunkID
	// map from chunk to chunk server address
	chunkLocation map[chunkID]*chunks.ChunkServiceClient
}

func (s MetaDataServer) GetFile(filename string) (*File, error) {
	var fileChunks []chunkID
	fileChunks, ok := s.fileMap[filename]
	if !ok {
		return nil, fmt.Errorf("file %s does not exist", filename)
	}
	if len(fileChunks) < 1 {
		return nil, fmt.Errorf("file %s contains no chunks", filename)
	}
	file := new(File)
	for _, id := range fileChunks {
		chunk, err := s.GetChunk(id)
		if err != nil {
			return nil, err
		}
		file.chunks = append(file.chunks, &chunk)
	}
	return file, nil
}

func (s MetaDataServer) GetChunk(id chunkID) (Chunck, error) {
	return Chunck{}, nil
}

const CHUNK_SIZE_BYTES = 1024

// implements fs.FS
var _ fs.FS = &TDFS{}

type ChunckServer struct {
	Port string
	ID   string
}

type TDFS struct {
	currentDir *directory
	rootDir    *directory
	server     MetaDataServer
	chunksize  int
}

// implements fs.FS
func (t *TDFS) Open(name string) (fs.File, error) {
	return &File{
		chunks: []*Chunck{},
	}, nil
}

func (m *TDFS) CreateFile(name, dir string, data []byte) error {
	d, err := m.rootDir.WalkTo(dir)
	if err != nil {
		return err
	}
	if _, ok := d.files[name]; ok {
		return fmt.Errorf("file %s already exists", name)
	}
	file := new(File)
	processed := 0
	for len(data) < processed {
		chunk := new(Chunck)
		chunk.ID = chunkID(uuid.New().String())
		for i := processed; i-processed < m.ChunkSize(); i++ {
			chunk.data = append(chunk.data, data[i])
		}
		file.chunks = append(file.chunks, chunk)
		// TODO:  send the chunk to one of the file servers
	}
	return nil
}

// ChunkSize returns the chunk size used
// in the file system
func (n *TDFS) ChunkSize() int {
	return n.chunksize
}

func (m *TDFS) ReadFile(name, dir string) ([]byte, error) {
	return []byte{}, nil
}
func (m *TDFS) DeleteFile(name, dir string) error {
	return nil
}
func (m *TDFS) CreateDir(dir string) error {
	return nil
}
func (m *TDFS) DeleteDir(dir string) error {
	return nil
}
func (m *TDFS) LS(dir string) error {
	return nil
}
func (m *TDFS) Grep(exp string) error {
	return nil
}

func (d *directory) WalkTo(p string) (*directory, error) {
	base := path.Base(p)
	rest := strings.TrimPrefix(p, base)
	for _, dir := range d.subDirs {
		if base == dir.name {
			return dir.WalkTo(rest)
		}
	}
	return nil, fmt.Errorf("directory does not exist: %s", base)
}

// creates a new TDFS filesystem with the given base directory
func NewTDFS() *TDFS {
	root := &directory{
		name:    "",
		prev:    nil,
		subDirs: []*directory{},
		files:   make(map[string]*File),
	}
	return &TDFS{
		currentDir: root,
		rootDir:    root,
		chunksize:  1024,
	}
}
