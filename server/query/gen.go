// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q             = new(Query)
	Auth          *auth
	AuthToken     *authToken
	Cert          *cert
	ChatGPTLog    *chatGPTLog
	ConfigBackup  *configBackup
	DnsCredential *dnsCredential
	Site          *site
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Auth = &Q.Auth
	AuthToken = &Q.AuthToken
	Cert = &Q.Cert
	ChatGPTLog = &Q.ChatGPTLog
	ConfigBackup = &Q.ConfigBackup
	DnsCredential = &Q.DnsCredential
	Site = &Q.Site
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:            db,
		Auth:          newAuth(db, opts...),
		AuthToken:     newAuthToken(db, opts...),
		Cert:          newCert(db, opts...),
		ChatGPTLog:    newChatGPTLog(db, opts...),
		ConfigBackup:  newConfigBackup(db, opts...),
		DnsCredential: newDnsCredential(db, opts...),
		Site:          newSite(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Auth          auth
	AuthToken     authToken
	Cert          cert
	ChatGPTLog    chatGPTLog
	ConfigBackup  configBackup
	DnsCredential dnsCredential
	Site          site
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		Auth:          q.Auth.clone(db),
		AuthToken:     q.AuthToken.clone(db),
		Cert:          q.Cert.clone(db),
		ChatGPTLog:    q.ChatGPTLog.clone(db),
		ConfigBackup:  q.ConfigBackup.clone(db),
		DnsCredential: q.DnsCredential.clone(db),
		Site:          q.Site.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		Auth:          q.Auth.replaceDB(db),
		AuthToken:     q.AuthToken.replaceDB(db),
		Cert:          q.Cert.replaceDB(db),
		ChatGPTLog:    q.ChatGPTLog.replaceDB(db),
		ConfigBackup:  q.ConfigBackup.replaceDB(db),
		DnsCredential: q.DnsCredential.replaceDB(db),
		Site:          q.Site.replaceDB(db),
	}
}

type queryCtx struct {
	Auth          *authDo
	AuthToken     *authTokenDo
	Cert          *certDo
	ChatGPTLog    *chatGPTLogDo
	ConfigBackup  *configBackupDo
	DnsCredential *dnsCredentialDo
	Site          *siteDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Auth:          q.Auth.WithContext(ctx),
		AuthToken:     q.AuthToken.WithContext(ctx),
		Cert:          q.Cert.WithContext(ctx),
		ChatGPTLog:    q.ChatGPTLog.WithContext(ctx),
		ConfigBackup:  q.ConfigBackup.WithContext(ctx),
		DnsCredential: q.DnsCredential.WithContext(ctx),
		Site:          q.Site.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
