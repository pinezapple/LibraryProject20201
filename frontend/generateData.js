module.exports = function(){
    var faker = require("faker")
    var _ = require("lodash")
    
    return {
// tài liệu
       documents: _.times(60, function(n){
        return {
            id:n,
            barcode : n,
            doc_name: faker.random.words(),
            doc_author: faker.name.findName(),
            doc_type: faker.commerce.color(),
            doc_publisher: faker.company.bs(10),
            doc_edition : faker.random.number(),
            doc_price : faker.commerce.price(),
            doc_status : faker.random.number(2),
            created_at: faker.date.recent(),
            updated_at: faker.date.recent(),
        }
       }),
// tài liệu bị hỏng
       doc_bad: _.times(60, function(n){
           return {
               id_sale: n,
               sold_by : faker.name.findName(),
               id_doc:n,
               barcode : faker.phone.phoneNumberFormat(),
               doc_name: faker.random.words(),
               doc_date : faker.date.recent(),
               doc_author: faker.name.findName(),
               librarian:faker.name.findName(),
               date: faker.date.recent(),
               fee: faker.commerce.price(),
               total: faker.commerce.price()
           }
       }),
// phiếu tài liệu thanh lý
       sale_form: _.times(60, function(n){
        return {
            id: n,
            id_librarian : faker.random.number(60),
            barcodes: [
                {barcode : faker.phone.phoneNumberFormat(),doc_price:faker.commerce.price()},
                {barcode : faker.phone.phoneNumberFormat(),doc_price:faker.commerce.price()}
            ],
            date: faker.date.recent(),
            
        }
    }),
// phiếu tài liệu thanh lý chi tiết
    detailed_sale_form: _.times(60, function(n){
        return {
            id: n,
            librarian_id : faker.random.number(60),
            librarian:faker.name.findName(),
            date: faker.date.recent(),
            barcodes: [
                {barcode : faker.phone.phoneNumberFormat(),doc_name: faker.random.words(),
                    doc_author: faker.name.findName(),doc_price:faker.commerce.price()},
                {barcode : faker.phone.phoneNumberFormat(),doc_name: faker.random.words(),
                    doc_author: faker.name.findName(),doc_price:faker.commerce.price()}
            ],
            date: faker.date.recent(),
            total : faker.commerce.price()
        }
    }),
// tài liệu đã mượn
       doc_returned: _.times(60, function(n){
        return {
            borrow_id: n,
            id_user:faker.random.number(60),
            name_user : faker.name.findName(),
            doc_status: faker.random.number({'min': 0,'max': 2 }),
            doc_date : faker.date.recent(),
            fee: faker.commerce.price(),
        }
    }),
// phiếu trả
       returned_form: _.times(60, function(n){
           return{
               id:n,
               reader_id:faker.random.number(60),
               reader: faker.name.findName(),
               librarian_id: faker.random.number(60),
               librarian:faker.name.findName(),
               barcode_update: [
                   {barcode: faker.phone.phoneNumberFormat(),id_doc:faker.random.number(60),doc_name: faker.random.words(),
                    doc_author: faker.name.findName(),doc_status :faker.random.number({'min': 0,'max': 2 }),doc_fine:faker.commerce.price()},
                    {barcode : faker.phone.phoneNumberFormat(),id_doc:faker.random.number(60),doc_name: faker.random.words(),
                    doc_author: faker.name.findName(),doc_status : faker.random.number({'min': 0,'max': 2 }),doc_fine:faker.commerce.price()},
                    {barcode : faker.phone.phoneNumberFormat(),id_doc:faker.random.number(60),doc_name: faker.random.words(),
                    doc_author: faker.name.findName(),doc_status : faker.random.number({'min': 0,'max': 2 }),doc_fine:faker.commerce.price()}
               ],
               duration: faker.random.number(7),
               borrow_status: faker.random.number({
                'min': 0,
                'max': 2
            }),
            fee: faker.random.number({'min': 1000,'max': 2000 })
           }
       }),
// phiếu mượn
       borrow_form: _.times(60, function(n){
        return{
            borrow_form_id:n,
            reader_id:faker.random.number(60),
            reader: faker.name.findName(),
            librarian_id: faker.random.number(60),
            barcodes: [],
            borrow_days: faker.random.number(7),
            borrow_status: faker.random.number({
             'min': 0,
             'max': 2
         })
        }
    }),
// Cập nhật phiếu mượn 
    updated_borrow_form: _.times(60, function(n){
        return{
            id:n,
            reader_id:faker.random.number(60),
            reader: faker.name.findName(),
            librarian_id: faker.random.number(60),
            barcodes: [
                {"barcode_id": faker.phone.phoneNumberFormat(),"barcode_status": faker.random.number({'min': 0,'max': 2 }),"fee":faker.commerce.price()},
                {"barcode_id": faker.phone.phoneNumberFormat(),"barcode_status": faker.random.number({'min': 0,'max': 2 }),"fee":faker.commerce.price()}
            ],
            borrow_days: faker.random.number(7),
            borrow_status: faker.random.number({
             'min': 0,
             'max': 2
         })
        }
    }),
    // vi pham detailed
    detailed_violation_form: _.times(60, function(n){
        return {
            id: n,
			user_name: faker.name.findName(),
			violate_count: faker.random.number(7),
            violated_books: [
                {borrow_ID : faker.phone.phoneNumberFormat(),violation_time: faker.date.recent(),
                    fee: faker.random.number({'min': 1000,'max': 20000 })},
                {borrow_ID : faker.phone.phoneNumberFormat(),violation_time: faker.date.recent(),
                    fee: faker.random.number({'min': 1000,'max': 20000 })},
            ],
            date: faker.date.recent(),
            total_fee : faker.commerce.price()
        };
    }),
// độc giả vi phạm
    violated_user: _.times(60, function (n) {
        return {
          user_id: faker.random.number(60),
          user_name: faker.name.findName(),
          violate_time: n,
		  total_fee: faker.random.number({'min': 1000,'max': 20000 })
        };
      }),
    }     
}