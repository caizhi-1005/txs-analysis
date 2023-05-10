package dbModels

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type TbTransaction struct {
	Id                int64     `orm:"column(id);pk"`
	BlockId           int64     `orm:"column(block_id);null" description:"区块号"`
	BlockHash         string    `orm:"column(block_hash);size(255);null" description:"区块hash"`
	TxHash            string    `orm:"column(tx_hash);size(255);null" description:"交易hash"`
	TxType            int       `orm:"column(tx_type);null" description:"交易类型 1：普通交易 2-合约交易"`
	TxTime            time.Time `orm:"column(tx_time);type(datetime);null" description:"交易时间"`
	TxIndex           int64     `orm:"column(tx_index);null" description:"交易索引"`
	ContractAddress   string    `orm:"column(contract_address);size(255);null" description:"合约地址"`
	From              string    `orm:"column(from);size(255);null" description:"发起地址"`
	To                string    `orm:"column(to);size(255);null" description:"到达地址"`
	Value             string    `orm:"column(value);size(255);null" description:"交易值"`
	Amount            float64   `orm:"column(amount);null;digits(64);decimals(18)" description:"交易金额"`
	TxFee             string    `orm:"column(tx_fee);size(255);null" description:"交易费"`
	TxStatus          int       `orm:"column(tx_status);null" description:"交易状态"`
	CumulativeGasUsed int64     `orm:"column(cumulative_gas_used);null" description:"CumulativeGasUsed"`
	GasUsed           int64     `orm:"column(gas_used);null" description:"使用gas"`
	GasLimit          int64     `orm:"column(gas_limit);null" description:"限制gas"`
	GasPrice          int64     `orm:"column(gas_price);null" description:"gas价格"`
	Nonce             int64     `orm:"column(nonce);null" description:"nonce号"`
	InputData         string    `orm:"column(input_data);null" description:"输入数据"`
}

type ShortTransaction struct {
	TxTime time.Time `json:"tx_time"`
	From   string    `json:"from"`
	To     string    `json:"to"`
	Amount float64   `json:"amount"`
}

func (t *TbTransaction) TableName() string {
	return "tb_transaction"
}

func init() {
	orm.RegisterModel(new(TbTransaction))
}

func GetWithdrawTxList() (Res []*TbTransaction, err error) {
	ormer := orm.NewOrm()
	//查询调用提现合约的交易 (to=0x871fcb6b836db1b5d6ee64901fb17245cd403e6d or input_data=0x86d1a69f)
	list := make([]*TbTransaction, 0)
	sql := "select * from tb_transaction where tx_status = 1 and input_data='0x86d1a69f' or `to`= '0x871fcb6b836db1b5d6ee64901fb17245cd403e6d'"
	_, err = ormer.Raw(sql).QueryRows(&list)
	if err != nil {
		return list, err
	}
	return list, nil
}

func WithdrawTxsListByAddress(address, inputData, contract string) (Res []*TbTransaction, err error) {
	cond := orm.NewCondition()
	list := make([]*TbTransaction, 0)
	query := orm.NewOrm().QueryTable(new(TbTransaction).TableName())
	cond = cond.And("input_data", inputData).Or("to", contract)
	if len(address) > 0 {
		cond = cond.And("from", address)
	}
	query = query.SetCond(cond)
	_, err = query.All(&list)
	if err != nil {
		return list, err
	}
	return list, nil
}

func AddressTxsList(address string) (Res []*TbTransaction, err error) {
	cond := orm.NewCondition()
	list := make([]*TbTransaction, 0)
	query := orm.NewOrm().QueryTable(new(TbTransaction).TableName())
	if len(address) > 0 {
		cond = cond.And("from", address)
	}
	query = query.SetCond(cond)
	_, err = query.All(&list)
	if err != nil {
		return list, err
	}
	return list, nil
}

func TransactionListGetByAddress(address string) ([]*ShortTransaction, error) {
	list := make([]*ShortTransaction, 0)
	query := orm.NewOrm().QueryTable(new(TbTransaction).TableName())
	cond := orm.NewCondition()
	cond.And("from", address).Or("to", address)
	query = query.OrderBy("-tx_time")
	_, err := query.All(list)
	if err != nil {
		return list, err
	}
	return list, nil
}
