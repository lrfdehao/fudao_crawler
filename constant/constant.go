package constant

const HOST_NAME = "https://fudao.qq.com"

const (
	// 获取所有年级科目信息
	API_GRADE_SUBJECT = HOST_NAME + "/cgi-proxy/course/grade_subject"

	// 获取年级、科目下的课程和课程包信息
	API_DISCOVER_SUBJECT = HOST_NAME + "/cgi-proxy/course/discover_subject"

	// 获取课程包下的课程信息
	API_GET_COURSE_PACKAGE = HOST_NAME + "/cgi-proxy/course/get_course_package_info"

	// 获取课程详情
	API_COURSE_DETAIL = HOST_NAME + "/cgi-proxy/courseLogic/courseStaticDetail"
)
