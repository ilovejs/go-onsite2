package glue

import (
	"onsite/models"
)

type UserStoreInterface interface {
	Create(*models.User) error
	GetByID(int) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	GetByUserName(string) (*models.User, error)
	Update(*models.User) error
	Delete(int) error
}

type ProjectStoreInterface interface {
	Create(*models.Project) error
	Read(int) (*models.Project, error)
	Update(*models.Project) error
	Delete(int) error

	//GetTradeSummary(uint)
	//GetTradeItems(uint, uint)            //assignTradeList(tid,pid)
	//GetTradeDetailByTradeId(uint, uint)  //GET addTradeItem
	//AddTradeItem(trade *model.TradeItem) //POST addTradeItem
	//GetTradeItemsByProjectId(uint)       //getSavedTradeItems(pid, tid)
	////todo:TradeItemsPartialSend
	//DeleteTradeItem(uint, string) //(id, reason)
	//UpdateTradeItems()            //EditTradeItem 2 methods
}

type ClaimStoreInterface interface {
	//Claim:
	//ProjectClaim(uint)              //Merge ctr with other ?
	//UpdateClaim(claim *model.Claim) //POST ProjectClaim
	//UpdateClaimByID(uint, uint)     //TradeItemsClaimEdit(pid, tid)
	//TradeTotalByProjectId(uint)     //GetSavedTradeTotal
	//TradeItemsClaimSummary(uint)    //getClaimSummaryByMonth
	//GetClaimByProjectId(uint)       //TradeItemsClaim
	//GetTradeDetail(uint)            //pid
	//TradeItemsClaimList(uint, uint) //GetTradeItems(uint, uint)
	//GetContractorClaimReport(uint)
	//SubmitClaim(uint, string, string)      //pid, month, claim number
	//GetClaimHistory(uint)                  //claimId
	//GetTradeItemListByLevel(uint, string)  //pid, level
	//GetProjectTradeItemByLevel(uint, uint) //pid, level
	//GetClaimReport(uint)
}
