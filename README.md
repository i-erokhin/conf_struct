conf_struct
===========

Golang library for build validated and filtered configuration structures from different sources.

[![Go Report Card](https://goreportcard.com/badge/github.com/i-erokhin/conf_struct)](https://goreportcard.com/report/github.com/i-erokhin/conf_struct)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/i-erokhin/conf_struct/blob/master/LICENSE)

Usage example
=============

Configuration structure definition:

```go
package config

type SmsRu struct {
    ApiId string
}

type Server struct {
    Addr      string
    URIprefix string
}

type DB struct {
    User                     string
    Password                 string
    ExternalAuth             bool
    Host                     string
    Port                     int
    ServiceName              string
    PoolMin                  int
    PoolMax                  int
    ManagerProcedurePassword string
}

type LDAP struct {
    Server            string
    PowerUserLogin    string
    PowerUserPassword string
    BaseDn            string
}

type Redis struct {
    Host    string
    PoolMax int
}

type Security struct {
    JWTSecret   string
    Supervisors []string
}

type Sentry struct {
    DSN string
}

type ORPC struct {
    Host string
    Port string
}

type TestUser struct {
    Barcode  *string
    Mobile   *string
    Login    *string
    Password *string
    Email    *string
}

type Conf struct {
    SmsRu       SmsRu
    Server      Server
    DB          DB
    LDAP        LDAP
    Redis       Redis
    Security    Security
    Sentry      Sentry
    ORPC        ORPC
    TestUser0   *TestUser
    TestUser1   *TestUser
    TestSeller  *TestUser
    TestManager *TestUser
}
```

Configuration structure initialization from environment variables with prefix ``PRJ_``, 
(``PRJ_DB_USER`` for example):

```go
package config

import (
    "fmt"

    "github.com/i-erokhin/conf_struct"
    "github.com/i-erokhin/conf_struct/sources/env"

    "project-specific-path/dft"
    "project-specific-path/core"
)

func Get() (*Conf, []error) {
    s := env.Source{Prefix: "AGS_"}
    b := conf_struct.Builder{}

    conf := Conf{
        SmsRu: SmsRu{
            ApiId: b.String("SMSRU_API_ID", s.Required),
        },
        Server: Server{
            Addr:      b.String("SERVER_ADDR", s.Default(dft.SERVER_ADDR)),
            URIprefix: b.String("SERVER_URI_PREFIX", s.Default(fmt.Sprintf("/api/v%d", core.ApiVersion))),
        },
        DB: DB{
            User:                     b.String("DB_USER", s.Required),
            Password:                 b.String("DB_PASSWORD", s.Required),
            ExternalAuth:             b.Bool("DB_EXTERNAL_AUTH", s.Default(dft.DB_EXTERNAL_AUTH)),
            Host:                     b.String("DB_HOST", s.Required),
            Port:                     b.Int("DB_PORT", s.Default(dft.DB_PORT)),
            ServiceName:              b.String("DB_SERVICE_NAME", s.Required),
            PoolMin:                  b.Int("DB_POOL_MIN", s.Default(dft.DB_POOL_MIN)),
            PoolMax:                  b.Int("DB_POOL_MAX", s.Default(dft.DB_POOL_MAX)),
            ManagerProcedurePassword: b.String("DB_MANAGER_PROCEDURE_PASSWORD", s.Required),
        },
        LDAP: LDAP{
            Server:            b.String("LDAP_SERVER", s.Required),
            PowerUserLogin:    b.String("LDAP_POWER_USER_LOGIN", s.Default(dft.LDAP_POWER_USER_LOGIN)),
            PowerUserPassword: b.String("LDAP_POWER_USER_PASSWORD", s.Required),
            BaseDn:            b.String("LDAP_BASE_DN", s.Default(dft.LDAP_BASE_DN)),
        },
        Redis: Redis{
            Host:    b.String("REDIS_HOST", s.Required),
            PoolMax: b.Int("REDIS_POOL_MAX", s.Default(dft.REDIS_POOL_MAX)),
        },
        Security: Security{
            JWTSecret:   b.String("SECURITY_JWT_SECRET", s.Required),
            Supervisors: b.StringArray("SECURITY_SUPERVISORS", s.Optional),
        },
        Sentry: Sentry{
            DSN: b.String("SENTRY_DSN", s.Required),
        },
        ORPC: ORPC{
            Host: b.String("ORPC_HOST", s.Default(dft.ORPC_HOST)),
            Port: b.String("ORPC_PORT", s.Default(dft.ORPC_PORT)),
        },
        TestUser0:   buildTestUser(&b, s, "TEST_USER_0_"),
        TestUser1:   buildTestUser(&b, s, "TEST_USER_1_"),
        TestSeller:  buildTestUser(&b, s, "TEST_SELLER_"),
        TestManager: buildTestUser(&b, s, "TEST_MANAGER_"),
    }

    return &conf, b.Errors
}

func buildTestUser(b *conf_struct.Builder, s conf_struct.Source, prefix string) *TestUser {
    u := TestUser{
        Barcode:  b.StringPointer(prefix+"BARCODE", s.Optional),
        Mobile:   b.StringPointer(prefix+"MOBILE", s.Optional),
        Login:    b.StringPointer(prefix+"LOGIN", s.Optional),
        Password: b.StringPointer(prefix+"PASSWORD", s.Optional),
        Email:    b.StringPointer(prefix+"EMAIL", s.Optional),
    }

    if u.Barcode == nil &&
        u.Mobile == nil &&
        u.Login == nil &&
        u.Password == nil &&
        u.Email == nil {
        return nil
    }
    
    return &u
}
```

License
=======

MIT

