package main

type Users struct {
	User_ID  int    `json:"user_id" gorm:"column:user_id;primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
}

type Subscriptions struct {
	ID_Payment    int    `json:"id_payment" gorm:"column:id_payment;primaryKey"`
	User_ID       int    `json:"user_id" gorm:"column:user_id"`
	Layanan_ID    int    `json:"layanan_id" gorm:"column:layanan_id"`
	Jenis_Payment string `json:"jenis_payment" gorm:"type:varchar(255);column:jenis_payment"`
	Active        bool   `json:"status_subscription" gorm:"column:active"`
}

type Services struct {
	Layanan_ID       int    `json:"layanan_id" gorm:"column:layanan_id;primaryKey;autoIncrement"`
	Nama_Layanan     string `json:"nama_layanan" gorm:"column:nama_layanan;type:varchar(255)"`
	Penyedia_Layanan string `json:"penyedia_layanan" gorm:"column:penyedia_layanan;type:varchar(255)"`
	Harga            int    `json:"harga" gorm:"column:harga"`
}
