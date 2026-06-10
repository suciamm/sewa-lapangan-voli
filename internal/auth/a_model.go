package auth

// Model layer untuk auth ditangani sepenuhnya oleh SQLC.
// File ini bisa digunakan untuk custom DB operation yang tidak bisa
// di-generate SQLC, misalnya raw query yang sangat dinamis.
//
// Untuk saat ini semua query auth sudah terdefinisi di:
//   - queries/users.sql
//   - queries/refresh_tokens.sql
//   - queries/email_verifications.sql
//
// Dan di-generate ke:
//   - db/users.sql.go
//   - db/refresh_tokens.sql.go
//   - db/email_verifications.sql.go
