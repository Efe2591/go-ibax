package model

type MintPoolTransferHistory struct {
	ID                       int64  `gorm:"not null"`
	Poolid                   int64  `gorm:"not null"`
	Keyid                    int64  `gorm:"not null"`
	Newkeyid                 int64  `gorm:"not null"`
	Status                   int64  `gorm:"not null"`
	MintTransfersHistoryHash []byte `gorm:"not null"`
	DateUpdated              int64  `gorm:"not null"`
	DateCreated              int64  `gorm:"not null"`
}

// TableName returns name of table
func (m MintPoolTransferHistory) TableName() string {
	return `1_mint_pool_transfer_history`
}
func (m *MintPoolTransferHistory) GetPool(keyid int64) (bool, error) {
	return isFound(DBConn.Where("keyid = ? and  status = ?", keyid, 1).Last(m))
}

// Get is retrieving model from database
func (m *MintPoolTransferHistory) GetPoolTransfer(poolid, keyid int64) (bool, error) {
	return isFound(DBConn.Where("poolid = ? and keyid = ? and  status = ?", poolid, keyid, 1).Last(m))
}
