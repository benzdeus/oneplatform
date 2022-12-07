package entities

type OneIDService interface {
	// Register()
	LoginPWD(username, password string) interface{}
	// LoginSharedToken()
	// GetSharedToken()
	// RefreshToken()
	// LoginOauthWithSharedToken()
	GetAccountData()
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
