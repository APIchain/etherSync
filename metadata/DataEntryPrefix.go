package metadata

type DataEntryPrefix byte

const (
	// DATA
	DATA_User              DataEntryPrefix = 0x00
	DATA_Contract          DataEntryPrefix = 0x01
	DATA_TokenTransfer     DataEntryPrefix = 0x02
	DATA_Transfer          DataEntryPrefix = 0x03
	DATA_All_TokenTransfer DataEntryPrefix = 0x04
	DATA_BLOCK_TIME        DataEntryPrefix = 0x05
	DATA_SC_LIST           DataEntryPrefix = 0x06
	DATA_TOKEN_LIST        DataEntryPrefix = 0x07
	DATA_USER_SC_LIST      DataEntryPrefix = 0x08

	//CONTEXT
	DATA_CONTEXT DataEntryPrefix = 0x10

	//SUMMARY
	DATA_User_Tx_Count        DataEntryPrefix = 0x20
	DATA_User_Token_Count     DataEntryPrefix = 0x21
	DATA_User_All_Token_Count DataEntryPrefix = 0x22

	//USER
	USER_CROWD_AMOUNT DataEntryPrefix = 0x30
)
