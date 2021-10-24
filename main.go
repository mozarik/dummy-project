package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

var ExcelMappingCell = map[string]string{
	"A1": "Outlet",
	"B1": "Nama Produk",
	"C1": "Kategori",
	"D1": "Deskripsi Produk",
	"E1": "Produk Favorite (Y/T)",
	"F1": "Tampilkan di menu (Y/T)",
	"G1": "Monitor Persediaan (Y/T)",
	"H1": "Stok Minimum",

	"I1": "Satuan #1",
	"J1": "Harga Beli Satuan #1",
	"K1": "Harga Jual Satuan #1",
	"L1": "SKU Satuan #1",
	"M1": "Minimum Transaksi Penjualan Satuan #1",

	"N1": "Satuan #2",
	"O1": "Rasio Satuan #2",
	"P1": "Harga Beli Satuan #2",
	"Q1": "Harga Jual Satuan #2",
	"R1": "SKU Satuan #2",
	"S1": "Minimum Transaksi Penjualan Satuan #2",

	"T1": "Satuan #3",
	"U1": "Rasio Satuan #3",
	"V1": "Harga Beli Satuan #3",
	"W1": "Harga Jual Satuan #3",
	"X1": "SKU Satuan #3",
	"Y1": "Minimum Transaksi Penjualan Satuan #3",

	"Z1":  "Satuan #4",
	"AA1": "Rasio Satuan #4",
	"AB1": "Harga Beli Satuan #4",
	"AC1": "Harga Jual Satuan #4",
	"AD1": "SKU Satuan #4",
	"AE1": "Minimum Transaksi Penjualan Satuan #4",

	"AF1": "Satuan #5",
	"AG1": "Rasio Satuan #5",
	"AH1": "Harga Beli Satuan #5",
	"AI1": "Harga Jual Satuan #5",
	"AJ1": "SKU Satuan #5",
	"AK1": "Minimum Transaksi Penjualan Satuan #5",

	"AL1": "Satuan Default",
	"AM1": "Satuan Penyimpanan",
	"AN1": "Satuan Pembelian",
	"AO1": "Ijinkan Kasir Ubah Harga Jual (Y/T)",
	"AP1": "Maksimal \"%\" dibawah harga jual",
}

func main() {
	xlxs, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		log.Print(err)
	}
	sheet1Name := "Sheet One"

	xlxs.SetSheetName(xlxs.GetSheetName(1), sheet1Name)

	// Setup Default Header
	for k, v := range ExcelMappingCell {
		xlxs.SetCellValue(sheet1Name, k, v)
	}
	err = xlxs.SaveAs("BookUpdate.xlsx")
	if err != nil {
		log.Print(err)
	}
}
