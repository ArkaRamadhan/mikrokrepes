package models

import (
	"encoding/json"
	"time"
)

type Memo struct {
	ID        uint       `gorm:"primaryKey"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
	Tanggal   *time.Time `json:"tanggal"`
	NoMemo    *string    `json:"no_memo"`
	Perihal   *string    `json:"perihal"`
	Pic       *string    `json:"pic"`
	CreateBy  string     `json:"create_by"`
}

// MarshalJSON menyesuaikan serialisasi JSON untuk struct Memo
func (i *Memo) MarshalJSON() ([]byte, error) {
	type Alias Memo
	if i.Tanggal == nil {
		return json.Marshal(&struct {
			Tanggal string `json:"tanggal"`
			*Alias
		}{
			Tanggal: "", // Atau nilai default lain yang Anda inginkan
			Alias:   (*Alias)(i),
		})
	} else {
		tanggalFormatted := i.Tanggal.Format("2006-01-02")
		return json.Marshal(&struct {
			Tanggal string `json:"tanggal"`
			*Alias
		}{
			Tanggal: tanggalFormatted,
			Alias:   (*Alias)(i),
		})
	}
}

type BeritaAcara struct {
	ID        uint       `gorm:"primaryKey"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
	NoSurat   *string    `json:"no_surat"`
	Tanggal   *time.Time `json:"tanggal"`
	Perihal   *string    `json:"perihal"`
	Pic       *string    `json:"pic"`
	CreateBy  string     `json:"create_by"`
}

func (i *BeritaAcara) MarshalJSON() ([]byte, error) {
	type Alias BeritaAcara
	if i.Tanggal == nil {
		return json.Marshal(&struct {
			Tanggal string `json:"tanggal"`
			*Alias
		}{
			Tanggal: "", // Atau nilai default lain yang Anda inginkan
			Alias:   (*Alias)(i),
		})
	} else {
		tanggalFormatted := i.Tanggal.Format("2006-01-02")
		return json.Marshal(&struct {
			Tanggal string `json:"tanggal"`
			*Alias
		}{
			Tanggal: tanggalFormatted,
			Alias:   (*Alias)(i),
		})
	}
}

type Surat struct {
	ID        uint       `gorm:"primaryKey"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
	NoSurat   *string    `json:"no_surat"`
	Tanggal   *time.Time `json:"tanggal"`
	Perihal   *string    `json:"perihal"`
	Pic       *string    `json:"pic"`
	CreateBy  string     `json:"create_by"`
}

func (i *Surat) MarshalJSON() ([]byte, error) {
	type Alias Surat
	if i.Tanggal == nil {
		return json.Marshal(&struct {
			Tanggal string `json:"tanggal"`
			*Alias
		}{
			Tanggal: "", // Atau nilai default lain yang Anda inginkan
			Alias:   (*Alias)(i),
		})
	} else {
		tanggalFormatted := i.Tanggal.Format("2006-01-02")
		return json.Marshal(&struct {
			Tanggal string `json:"tanggal"`
			*Alias
		}{
			Tanggal: tanggalFormatted,
			Alias:   (*Alias)(i),
		})
	}
}

type Sk struct {
	ID        uint       `gorm:"primaryKey"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
	NoSurat   *string    `json:"no_surat"`
	Tanggal   *time.Time `json:"tanggal"`
	Perihal   *string    `json:"perihal"`
	Pic       *string    `json:"pic"`
	CreateBy  string     `json:"create_by"`
}

func (i *Sk) MarshalJSON() ([]byte, error) {
	type Alias Sk
	if i.Tanggal == nil {
		return json.Marshal(&struct {
			Tanggal string `json:"tanggal"`
			*Alias
		}{
			Tanggal: "", // Atau nilai default lain yang Anda inginkan
			Alias:   (*Alias)(i),
		})
	} else {
		tanggalFormatted := i.Tanggal.Format("2006-01-02")
		return json.Marshal(&struct {
			Tanggal string `json:"tanggal"`
			*Alias
		}{
			Tanggal: tanggalFormatted,
			Alias:   (*Alias)(i),
		})
	}
}
