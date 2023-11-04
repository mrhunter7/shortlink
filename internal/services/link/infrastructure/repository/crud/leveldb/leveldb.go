package leveldb

import (
	"context"

	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/query"
)

// Store implementation of db interface
type Store struct {
	client *leveldb.DB
}

// New store
func New(_ context.Context, store db.DB) (*Store, error) {
	s := &Store{
		client: store.GetConn().(*leveldb.DB),
	}

	return s, nil
}

// Add - add
func (l *Store) Add(_ context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	payload, err := protojson.Marshal(source)
	if err != nil {
		return nil, err
	}

	err = l.client.Put([]byte(source.GetHash()), payload, nil)
	if err != nil {
		return nil, err
	}

	return source, nil
}

// Get - get
func (l *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	value, err := l.client.Get([]byte(id), nil)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}}
	}

	var response v1.Link
	err = protojson.Unmarshal(value, &response)
	if err != nil {
		return nil, err
	}

	if response.GetUrl() == "" {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}}
	}

	return &response, nil
}

// List - list
func (l *Store) List(_ context.Context, _ *query.Filter) (*v1.Links, error) {
	links := &v1.Links{
		Link: []*v1.Link{},
	}
	iterator := l.client.NewIterator(nil, nil)

	for iterator.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		value := iterator.Value()

		var response v1.Link
		err := protojson.Unmarshal(value, &response)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}}
		}

		links.Link = append(links.GetLink(), &response)
	}

	iterator.Release()
	err := iterator.Error()
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}}
	}

	return links, nil
}

// Update - update
func (l *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete - delete
func (l *Store) Delete(ctx context.Context, id string) error {
	err := l.client.Delete([]byte(id), nil)
	return err
}
