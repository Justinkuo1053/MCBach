package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	sqldriver "github.com/go-sql-driver/mysql" // 指定別名
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RegisterCustomTLSConfig() {
	// 建立一個新的根憑證池
	rootCertPool := x509.NewCertPool()

	// 讀取 CA 憑證
	pemPath := "E:/CS學習 (E曹)/GO 專案/MCBach-main(go language)/ca.pem"
	pem, err := os.ReadFile(pemPath)
	if err != nil {
		log.Fatalf("無法讀取 CA 憑證: %v", err)
	}
	// 將 CA 憑證加入根憑證池
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatalf("無法載入 CA 憑證")
	}

	// 建立自定義的 TLS 配置
	tlsConfig := &tls.Config{
		RootCAs: rootCertPool,
	}

	// 註冊自定義的 TLS 配置
	err = sqldriver.RegisterTLSConfig("custom", tlsConfig)
	if err != nil {
		log.Fatalf("無法註冊自定義 TLS 配置: %v", err)
	}

	log.Println("自定義 TLS 配置已成功註冊")
}

func ConnectDB() *gorm.DB {
	// 明確指定 .env 檔案的路徑
	envPath := "E:/CS學習 (E曹)/GO 專案/MCBach-main(go language)/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf(".env 檔案未找到，將使用系統環境變數: %v", err)
	}

	// 註冊自定義的 TLS 配置
	RegisterCustomTLSConfig()

	// 從環境變數中讀取資料庫連線字串
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("環境變數 DATABASE_URL 未設定")
	}

	// 在 DSN 中加入自定義的 TLS 配置
	dsnWithTLS := fmt.Sprintf("%s&tls=custom", dsn)

	// 使用 MySQL 作為驅動
	db, err := gorm.Open(mysql.Open(dsnWithTLS), &gorm.Config{})
	if err != nil {
		log.Fatalf("無法連接到資料庫: %v", err)
	}

	log.Println("成功連接到 MySQL 資料庫")
	return db
}

