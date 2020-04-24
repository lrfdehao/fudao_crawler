package model

type GradeSubjResp struct {
	Result struct {
		GradeSubjects []struct {
			Grade   int64   `json:"grade"`
			Subject []int64 `json:"subject"`
		} `json:"grade_subjects"`
	} `json:"result"`
	Retcode int `json:"retcode"`
}

type DiscoverSubjResp struct {
	Result struct {
		Grade         int64 `json:"grade"`
		Retcode       int64 `json:"retcode"`
		SpeCourseList struct {
			Data  []SpeCourse `json:"data"`
			Page  int64       `json:"page"`
			Size  int64       `json:"size"`
			Total int64       `json:"total"`
		} `json:"spe_course_list"`
		SysCoursePkgList []struct {
			CidList           string `json:"cid_list"`
			CourseBgtime      int64  `json:"course_bgtime"`
			CourseEndtime     int64  `json:"course_endtime"`
			CourseMaxPrice    int64  `json:"course_max_price"`
			CourseMinPrice    int64  `json:"course_min_price"`
			CourseSignEndtime int64  `json:"course_sign_endtime"`
			DiscountPrice     int64  `json:"discount_price"`
			Season            int64  `json:"season"`
			SoldCount         int64  `json:"sold_count"`
			SubjectPackageID  string `json:"subject_package_id"`
			Title             string `json:"title"`
		} `json:"sys_course_pkg_list"`
	} `json:"result"`
	Retcode int64 `json:"retcode"`
}

type SpeCourse struct {
	AfAmount  int64 `json:"af_amount"`
	Applied   int64 `json:"applied"`
	ApplyNum  int64 `json:"apply_num"`
	Cid       int64 `json:"cid"`
	ClassInfo struct {
		ClassID int64 `json:"class_id"`
	} `json:"class_info"`
	ClassType    int64 `json:"class_type"`
	CourseLabels []struct {
		LabelOptionID   int64  `json:"label_option_id"`
		LabelOptionName string `json:"label_option_name"`
		LabelOptionType int64  `json:"label_option_type"`
	} `json:"course_labels"`
	CoverURL        string `json:"cover_url"`
	CoverURLColor   string `json:"cover_url_color"`
	DisBgtime       int64  `json:"dis_bgtime"`
	DisEdtime       int64  `json:"dis_edtime"`
	FirstSubBgtime  int64  `json:"first_sub_bgtime"`
	FirstSubEndtime int64  `json:"first_sub_endtime"`
	Grade           int64  `json:"grade"`
	HasDiscount     int64  `json:"has_discount"`
	HintLogo        string `json:"hint_logo"`
	LastSubBgtime   int64  `json:"last_sub_bgtime"`
	LastSubEndtime  int64  `json:"last_sub_endtime"`
	Level           int64  `json:"level"`
	Name            string `json:"name"`
	Payid           int64  `json:"payid"`
	Pinyin          string `json:"pinyin"`
	PreAmount       int64  `json:"pre_amount"`
	Recordtime      int64  `json:"recordtime"`
	RenewType       int64  `json:"renew_type"`
	Showid          int64  `json:"showid"`
	SignBgtime      int64  `json:"sign_bgtime"`
	SignEndTime     int64  `json:"sign_end_time"`
	SignMax         int64  `json:"sign_max"`
	Status          int64  `json:"status"`
	StudentTotal    int64  `json:"student_total"`
	Subject         int64  `json:"subject"`
	TeList          []struct {
		CoverSmallURL string `json:"cover_small_url"`
		CoverURL      string `json:"cover_url"`
		Introduce     string `json:"introduce"`
		Name          string `json:"name"`
		PlatfromRole  int64  `json:"platfrom_role"`
	} `json:"te_list"`
	TimePlan string `json:"time_plan"`
}

type PackageCourseInfo struct {
	Result struct {
		Courses          []SpeCourse `json:"courses"`
		Retcode          int64       `json:"retcode"`
		SysCoursePkgInfo struct {
			CidList           string `json:"cid_list"`
			CourseBgtime      int64  `json:"course_bgtime"`
			CourseEndtime     int64  `json:"course_endtime"`
			CourseMaxPrice    int64  `json:"course_max_price"`
			CourseMinPrice    int64  `json:"course_min_price"`
			CourseSignEndtime int64  `json:"course_sign_endtime"`
			DiscountPrice     int64  `json:"discount_price"`
			Season            int64  `json:"season"`
			SoldCount         int64  `json:"sold_count"`
			SubjectPackageID  string `json:"subject_package_id"`
			Title             string `json:"title"`
		} `json:"sys_course_pkg_info"`
	} `json:"result"`
	Retcode int64 `json:"retcode"`
}

type CourseDetailResp struct {
	Result struct {
		BgColor        string `json:"bg_color"`
		BgTime         int64  `json:"bg_time"`
		Cid            int64  `json:"cid"`
		CourseLogo     string `json:"course_logo"`
		CourseTimeDesc string `json:"course_time_desc"`
		CoverURL       string `json:"cover_url"`
		Desc           string `json:"desc"`
		Directory      []struct {
			Bgtime        int64    `json:"bgtime"`
			CsRoomType    int64    `json:"cs_room_type"`
			CsType        int64    `json:"cs_type"`
			Csid          int64    `json:"csid"`
			Endtime       int64    `json:"endtime"`
			ExamDuration  int64    `json:"exam_duration"`
			Tid           int64    `json:"tid"`
			Title         string   `json:"title"`
			Tname         string   `json:"tname"`
			VideoRoomInfo struct{} `json:"video_room_info"`
		} `json:"directory"`
		EndTime      int64   `json:"end_time"`
		FirstTime    int64   `json:"first_time"`
		Grade        string  `json:"grade"`
		IosCheck     int64   `json:"ios_check"`
		IsFlowCourse int64   `json:"is_flow_course"`
		IsSmallTerm  int64   `json:"is_small_term"`
		Knowledge    string  `json:"knowledge"`
		Level        int64   `json:"level"`
		Name         string  `json:"name"`
		PhoneRequire int64   `json:"phone_require"`
		Price        int64   `json:"price"`
		Reshelf      int64   `json:"reshelf"`
		ServiceQq    []int64 `json:"service_qq"`
		SignBgtime   int64   `json:"sign_bgtime"`
		SignEndtime  int64   `json:"sign_endtime"`
		Subject      int64   `json:"subject"`
		Suitable     []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"suitable"`
		Target  []string `json:"target"`
		Teacher []struct {
			Content  string `json:"content"`
			Desc     string `json:"desc"`
			Exp      string `json:"exp"`
			GoodRate string `json:"good_rate"`
			Label    string `json:"label"`
			Name     string `json:"name"`
			Pic      string `json:"pic"`
			Tid      int64  `json:"tid"`
		} `json:"teacher"`
		TermID int64  `json:"term_id"`
		Type   int64  `json:"type"`
		Vid    string `json:"vid"`
	} `json:"result"`
	Retcode int64 `json:"retcode"`
}
