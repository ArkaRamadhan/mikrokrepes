package models

import (
	"encoding/json"
	"time"
)

type Project struct {
	ID              uint       `gorm:"primaryKey"`
	CreatedAt       *time.Time `gorm:"autoCreateTime"`
	UpdatedAt       *time.Time `gorm:"autoUpdateTime"`
	KodeProject     *string    `json:"kode_project"`
	JenisPengadaan  *string    `json:"jenis_pengadaan"`
	NamaPengadaan   *string    `json:"nama_pengadaan"`
	DivInisiasi     *string    `json:"div_inisiasi"`
	Bulan           *time.Time `json:"bulan"`
	SumberPendanaan *string    `json:"sumber_pendanaan"`
	Anggaran        *string    `json:"anggaran"`
	NoIzin          *string    `json:"no_izin"`
	TanggalIzin     *time.Time `json:"tanggal_izin"`
	TanggalTor      *time.Time `json:"tanggal_tor"`
	Pic             *string    `json:"pic"`
	Group           *string    `json:"group"`
	InfraType       *string    `json:"infra_type"`
	BudgetType      *string    `json:"budget_type"`
	Type            *string    `json:"type"`
	CreateBy        string     `json:"create_by"`
}

func (p *Project) MarshalJSON() ([]byte, error) {
	type Alias Project
	var tanggalIzinFormatted, tanggalTorFormatted, bulanFormatted string

	// Cek TanggalIzin
	if p.TanggalIzin == nil {
		tanggalIzinFormatted = ""
	} else {
		tanggalIzinFormatted = p.TanggalIzin.Format("2006-01-02")
	}

	// Cek TanggalTor
	if p.TanggalTor == nil {
		tanggalTorFormatted = ""
	} else {
		tanggalTorFormatted = p.TanggalTor.Format("2006-01-02")
	}

	// Cek Bulan
	if p.Bulan == nil {
		bulanFormatted = ""
	} else {
		bulanFormatted = p.Bulan.Format("01/06")
	}

	return json.Marshal(&struct {
		TanggalIzin string `json:"tanggal_izin"`
		TanggalTor  string `json:"tanggal_tor"`
		Bulan       string `json:"bulan"`
		*Alias
	}{
		TanggalIzin: tanggalIzinFormatted,
		TanggalTor:  tanggalTorFormatted,
		Bulan:       bulanFormatted,
		Alias:       (*Alias)(p),
	})
}
