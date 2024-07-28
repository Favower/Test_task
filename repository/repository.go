package repository

import (
    "context"
    "github.com/jackc/pgx/v4"
    "app/model"
)

type Repository struct {
    DB *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
    return &Repository{DB: db}
}

func (r *Repository) SaveMessage(message *model.Message) error {
    query := `INSERT INTO messages (content, processed, created_at) VALUES ($1, $2, $3) RETURNING id`
    return r.DB.QueryRow(context.Background(), query, message.Content, message.Processed, message.CreatedAt).Scan(&message.ID)
}

func (r *Repository) GetProcessedMessages() ([]model.Message, error) {
    query := `SELECT id, content, processed, created_at FROM messages WHERE processed = TRUE`
    rows, err := r.DB.Query(context.Background(), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var messages []model.Message
    for rows.Next() {
        var message model.Message
        if err := rows.Scan(&message.ID, &message.Content, &message.Processed, &message.CreatedAt); err != nil {
            return nil, err
        }
        messages = append(messages, message)
    }

    return messages, nil
}
