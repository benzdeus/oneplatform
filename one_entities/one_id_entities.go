package one_entities

import "github.com/benzdeus/oneplatform/env"

type OneIDService interface {
	// Method สำหรับสมัครสมาชิก
	Register(RequestOneIDRegister) ResponseOneIDRegister

	// Method สำหรับ Login เพื่อดึง Access Token
	//
	// username คือ username ของ one id
	//
	// password คือ password ของ one id
	LoginPWD(username, password string) ResponseOneIDLoginPWD

	// LoginSharedToken()
	// GetSharedToken()
	// RefreshToken()
	// LoginOauthWithSharedToken()

	// Method สำหรับดึงข้อมูล user โดยใช้ access token
	//
	// accessToken คือ access token ที่ได้มาจาก method LoginPWD
	GetAccountData(accessToken string) ResponseOneIDGetAccountData

	GetAccessToken(sharedToken string) ResponseONEIDGetAccessToken

	PostCheckUsernameAndMobile(username, mobile_no string) ResponseOneIDCheckUsernameAndMobile

	PostDeleteEmail(email, accessToken string) ResponseOneIDDeleteEmail

	GetCheckEmail(email string) ResponseOneIDCheckEmail

	PostCheckOneMail(onemail string) ResponseOneIDCheckOneMail

	PostCheckUsernameAndIdCard(username, id_card string) ResponseOneIdCheckUsernameAndIdCard

	GetCheckId(idCard string) ResponseOneIdCheckId

	GetAccountDetialByIdNo(id_no, access_token string) ResponseOneIdAccountDetailByIdNo
}

type RequestOneIDRegister struct {
	AccountTitleTh  string `json:"account_title_th"`
	FirstNameTh     string `json:"first_name_th"`
	LastNameTh      string `json:"last_name_th"`
	AccountTitleEng string `json:"account_title_eng"`
	FirstNameEng    string `json:"first_name_eng"`
	LastNameEng     string `json:"last_name_eng"`
	IdCardType      string `json:"id_card_type"`
	IdCardNum       string `json:"id_card_num"`
	Email           string `json:"email"`
	MobileNo        string `json:"mobile_no"`
	BirthDate       string `json:"birth_date"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ClientId        string `json:"clientId"`
	SecretKey       string `json:"secretKey"`
	RefCode         string `json:"ref_code"`
}

func (r *RequestOneIDRegister) SetKey() {
	(*r).ClientId = env.Options.ClientID
	(*r).SecretKey = env.Options.ClientSecret
	(*r).RefCode = env.Options.RefCode
}

func (r *RequestOneIDRegister) GetKey() (string, string) {

	return (*r).ClientId, (*r).SecretKey
}

type ResponseOneIDRegister struct {
	Result string `json:"result"`
	Data   struct {
		AccountID string `json:"accountID"`
		Email     string `json:"email"`
		OneChat   string `json:"one_chat"`
		OneBox    string `json:"one_box"`
	} `json:"data"`
	ErrorMessage map[string][]string `json:"errorMessage"`
	Code         int                 `json:"code"`
}

type ResponseOneIDLoginPWD struct {
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	ExpirationDate string `json:"expiration_date"`
	AccountID      string `json:"account_id"`
	Result         string `json:"result"`
	UserName       string `json:"username"`
	ResponseOneIDFail
}

type ResponseOneIDFail struct {
	Result       string      `json:"result"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
	ResponseCode int         `json:"responseCode"`
}

type ResponseONEIDGetAccessToken struct {
	TokenType      string  `json:"token_type"`
	ExpiresIn      int     `json:"expires_in"`
	AccessToken    string  `json:"access_token"`
	RefreshToken   *string `json:"refresh_token"`
	ExpirationDate string  `json:"expiration_date"`
	AccountID      string  `json:"account_id"`
	Result         string  `json:"result"`
	Username       string  `json:"username"`
}

type ResponseOneIDGetAccountData struct {
	Id                  string  `json:"id"`
	FirstNameTh         string  `json:"first_name_th"`
	MiddleNameTh        *string `json:"middle_name_th"`
	LastNameTh          string  `json:"last_name_th"`
	FirstNameEng        *string `json:"first_name_eng"`
	MiddleNameEng       *string `json:"middle_name_eng"`
	LastNameEng         *string `json:"last_name_eng"`
	SpecialTitleNameTh  *string `json:"special_title_name_th"`
	AccountTitleTh      string  `json:"account_title_th"`
	SpecialTitleNameEng *string `json:"special_title_name_eng"`
	AccountTitleEng     string  `json:"account_title_eng"`
	IdCardType          string  `json:"id_card_type"`
	IdCardNum           string  `json:"id_card_num"`
	HashIdCardNum       string  `json:"hash_id_card_num"`
	AccountCategory     string  `json:"account_category"`
	AccountSubCategory  string  `json:"account_sub_category"`
	ThaiEmail           string  `json:"thai_email"`
	ThaiEmail2          *string `json:"thai_email2"`
	ThaiEmail3          *string `json:"thai_email3"`
	StatusCd            string  `json:"status_cd"`
	BirthDate           string  `json:"birth_date"`
	StatusDt            string  `json:"status_dt"`
	RegisterDt          string  `json:"register_dt"`
	AddressId           *int    `json:"address_id"`
	CreatedAt           string  `json:"created_at"`
	CreatedBy           string  `json:"created_by"`
	UpdatedAt           string  `json:"updated_at"`
	UpdatedBy           string  `json:"updated_by"`
	Reason              *string `json:"reason"`
	TelNo               *string `json:"tel_no"`
	NameOnDocumentTh    *string `json:"name_on_document_th"`
	NameOnDocumentEng   *string `json:"name_on_document_eng"`
	BlockchainFlg       string  `json:"blockchain_flg"`
	Mobile              []struct {
		Id        string      `json:"id"`
		MobileNo  string      `json:"mobile_no"`
		CreatedAt string      `json:"created_at"`
		CreatedBy string      `json:"created_by"`
		UpdatedAt string      `json:"updated_at"`
		UpdatedBy string      `json:"updated_by"`
		DeletedAt interface{} `json:"deleted_at"`
		Pivot     struct {
			AccountId             string      `json:"account_id"`
			MobileId              string      `json:"mobile_id"`
			CreatedAt             string      `json:"created_at"`
			UpdatedAt             string      `json:"updated_at"`
			StatusCd              string      `json:"status_cd"`
			PrimaryFlg            string      `json:"primary_flg"`
			MobileConfirmFlg      string      `json:"mobile_confirm_flg"`
			MobileConfirmDt       *string     `json:"mobile_confirm_dt"`
			MobileLoginConfirmFlg string      `json:"mobile_login_confirm_flg"`
			MobileLoginConfirmDt  interface{} `json:"mobile_login_confirm_dt"`
			Type                  interface{} `json:"type"`
		} `json:"pivot"`
	} `json:"mobile"`
	Email []struct {
		Id        string      `json:"id"`
		Email     string      `json:"email"`
		CreatedAt string      `json:"created_at"`
		CreatedBy string      `json:"created_by"`
		UpdatedAt string      `json:"updated_at"`
		UpdatedBy string      `json:"updated_by"`
		DeletedAt interface{} `json:"deleted_at"`
		Pivot     struct {
			AccountId            string      `json:"account_id"`
			EmailId              string      `json:"email_id"`
			CreatedAt            string      `json:"created_at"`
			UpdatedAt            string      `json:"updated_at"`
			StatusCd             string      `json:"status_cd"`
			PrimaryFlg           string      `json:"primary_flg"`
			EmailConfirmFlg      string      `json:"email_confirm_flg"`
			EmailConfirmDt       interface{} `json:"email_confirm_dt"`
			EmailLoginConfirmFlg string      `json:"email_login_confirm_flg"`
			EmailLoginConfirmDt  interface{} `json:"email_login_confirm_dt"`
		} `json:"pivot"`
	} `json:"email"`
	Address []struct {
		Id           string      `json:"id"`
		HouseNo      string      `json:"house_no"`
		MooBan       *string     `json:"moo_ban"`
		BuildingName *string     `json:"building_name"`
		Street       *string     `json:"street"`
		Soi          *string     `json:"soi"`
		ZipcodeId    string      `json:"zipcode_id"`
		CreatedAt    string      `json:"created_at"`
		CreatedBy    string      `json:"created_by"`
		UpdatedAt    string      `json:"updated_at"`
		UpdatedBy    string      `json:"updated_by"`
		DeletedAt    interface{} `json:"deleted_at"`
		RoomNo       *string     `json:"room_no"`
		Floor        *string     `json:"floor"`
		FaxNumber    interface{} `json:"fax_number"`
		Type         string      `json:"type"`
		HouseCode    interface{} `json:"house_code"`
		MooNo        *string     `json:"moo_no"`
		DataDetail   interface{} `json:"data_detail"`
		CountryId    string      `json:"country_id"`
		Yaek         *string     `json:"yaek"`
		Pivot        struct {
			AccountId  string      `json:"account_id"`
			AddressId  string      `json:"address_id"`
			CreatedAt  string      `json:"created_at"`
			UpdatedAt  string      `json:"updated_at"`
			StatusCd   string      `json:"status_cd"`
			PrimaryFlg string      `json:"primary_flg"`
			Type       interface{} `json:"type"`
		} `json:"pivot"`
		Country struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			Alpha2    string `json:"alpha_2"`
			Alpha3    string `json:"alpha_3"`
			Numeric   string `json:"numeric"`
			Iso       string `json:"iso"`
			CreatedAt string `json:"created_at"`
			CreatedBy string `json:"created_by"`
			UpdatedAt string `json:"updated_at"`
			UpdatedBy string `json:"updated_by"`
		} `json:"country"`
		Zipcode struct {
			Id              string `json:"id"`
			Amphoe          string `json:"amphoe"`
			Tambon          string `json:"tambon"`
			Province        string `json:"province"`
			Zipcode         string `json:"zipcode"`
			SubdistrictCode string `json:"subdistrict_code"`
			DistrictCode    string `json:"district_code"`
			ProvinceCode    string `json:"province_code"`
			CreatedAt       string `json:"created_at"`
			CreatedBy       string `json:"created_by"`
			UpdatedAt       string `json:"updated_at"`
			UpdatedBy       string `json:"updated_by"`
		} `json:"zipcode"`
	} `json:"address"`
	AccountAttribute []interface{} `json:"account_attribute"`
	Status           string        `json:"status"`
	LastUpdate       string        `json:"last_update"`
	ResponseOneIDFail
}

type ResponseOneIDCheckUsernameAndMobile struct {
	ResponseOneIDFail
}

type ResponseOneIDDeleteEmail struct {
	ResponseOneIDFail
}

type ResponseOneIDCheckEmail struct {
	ResponseOneIDFail
}

type ResponseOneIDCheckOneMail struct {
	ResponseOneIDFail
}

type ResponseOneIdCheckUsernameAndIdCard struct {
	ResponseOneIDFail
}

type ResponseOneIdCheckId struct {
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
	Message      string      `json:"message"`
	ResponseCode int         `json:"responseCode"`
}

type ResponseOneIdAccountDetailByIdNo struct {
	Result string `json:"result"`
	Data   []struct {
		Id              string        `json:"id"`
		FirstNameTh     string        `json:"first_name_th"`
		LastNameTh      string        `json:"last_name_th"`
		FirstNameEng    string        `json:"first_name_eng"`
		LastNameEng     string        `json:"last_name_eng"`
		AccountTitleTh  string        `json:"account_title_th"`
		AccountTitleEng string        `json:"account_title_eng"`
		AccountCategory string        `json:"account_category"`
		IdCardType      string        `json:"id_card_type"`
		IdCardNum       string        `json:"id_card_num"`
		StatusCd        string        `json:"status_cd"`
		Address         []interface{} `json:"address"`
	}
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
	ResponseCode int         `json:"responseCode"`
}

type ResponseOneIdCheckLoaLevel1To3 struct {
	Result string `json:"result"`
	Data   struct {
		LoaLevel string `json:"loa_level"`
		Loa1     string `json:"loa1"`
		Loa2     string `json:"loa2"`
		Loa3     string `json:"loa3"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
}

type ResponseOneIdAddAddress struct {
	ResponseOneIDFail
}
