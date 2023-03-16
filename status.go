package thailandposttracking

import "net/http"

const (
	ItemStatusAll = "all"
	ItemStatus101 = "101" // เตรียมการฝากส่ง
	ItemStatus102 = "102" // รับฝากผ่านตัวแทน
	ItemStatus103 = "103" // รับฝาก
	ItemStatus104 = "104" // ผู้ฝากส่งขอถอนคืน / ยกเลิกการรับฝาก
	ItemStatus201 = "201" // อยู่ระหว่างการขนส่ง
	ItemStatus202 = "202" // ดำเนินพิธีการศุลกากร
	ItemStatus203 = "203" // ส่งคืนต้นทาง
	ItemStatus204 = "204" // ถึงที่ทำการแลกเปลี่ยนระหว่างประเทศขาออก
	ItemStatus205 = "205" // ถึงที่ทำการแลกเปลี่ยนระหว่างประเทศขาเข้า
	ItemStatus206 = "206" // ถึงที่ทำการไปรษณีย์
	ItemStatus208 = "208" // ส่งออกจากที่ทำการแลกเปลี่ยนระหว่างประเทศขาออก
	ItemStatus209 = "209" // ยกเลิกการส่งออก
	ItemStatus210 = "210" // ยกเลิกการนำเข้า
	ItemStatus211 = "211" // รับเข้า ณ ศูนย์คัดแยก
	ItemStatus212 = "212" // ส่งมอบให้สายการบิน
	ItemStatus213 = "213" // สายการบินรับมอบ
	ItemStatus301 = "301" // อยู่ระหว่างการนำจ่าย
	ItemStatus302 = "302" // นำจ่าย ณ จุดรับสิ่งของ
	ItemStatus401 = "401" // นำจ่ายไม่สำเร็จ
	ItemStatus402 = "402" // ปิดประกาศ ณ ที่ทำการรับฝาก : กรุณาติดต่อ THP contact center 1545
	ItemStatus501 = "501" // นำจ่ายสำเร็จ
	ItemStatus901 = "901" // โอนเงินให้ผู้ขายเรียบร้อยแล้ว
)

const (
// HTTP 200 OK บันทึกสำเร็จ(รวมบันทึกได้ทั้งหมดและบางส่วน)ให้ตรวจสอบ received_record ควบคู่ หากระบบบันทึกข้อมูลได้บางส่วน จำนาน received_record จะไม่เท่ากับจำนวน locations ที่ส่ง
// HTTP 400 Bad Request รูปแบบข้อมูลที่ไม่ถูกต้อง ไม่สามารถบันทึกได้ ระบบจะแนบรายละเอียดข้อผิดพลาดกลับมาด้วย โดยมีรูปแบบดังนี้
// HTTP 500 Internal Server Error บันทึกไม่สำเร็จ(เกิดจากปัญหาภายในระบบ)
// HTTP 403 Forbidden (เกิดในขั้นตอน Authentication) ไม่ผ่านการอนุญาติ
// HTTP 401 Unauthorized Login Credentials ไม่ถูกต้อง
// HTTP 429 Too Many Requests	จำนวนครั้งในการ Request มากเกินกว่าที่กำหนด
)

func thailandpostStatusTHText(code int) string {
	switch code {
	case http.StatusOK:
		return "บันทึกสำเร็จ"
	case http.StatusBadRequest:
		return "รูปแบบข้อมูลที่ไม่ถูกต้อง"
	case http.StatusInternalServerError:
		return "บันทึกไม่สำเร็จ(เกิดจากปัญหาภายในระบบ)"
	case http.StatusForbidden:
		return "ไม่ผ่านการอนุญาติ"
	case http.StatusUnauthorized:
		return "Login Credentials ไม่ถูกต้อง"
	case http.StatusTooManyRequests:
		return "จำนวนครั้งในการ Request มากเกินกว่าที่กำหนด"
	}
	return http.StatusText(code)
}
