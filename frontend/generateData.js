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
            date: faker.date.recent()
        }
    }),
// tài liệu đã mượn
       doc_returned: _.times(60, function(n){
        return {
            borrow_id: n,
            id_user:faker.random.number(60),
            name_user : faker.name.findName(),
            doc_status: n,
            doc_date : faker.date.recent(),
            fee: faker.commerce.price(),
        }
    }),
// phiếu trả
       returned_form: _.times(60, function(n){
           return{
               id:n,
               id_reader:faker.random.number(60),
               reader: faker.name.findName(),
               id_librarian: faker.random.number(60),
               librarian:faker.name.findName(),
               docs: [
                   {barcode : faker.phone.phoneNumberFormat(),id_doc:faker.random.number(60),doc_name: faker.random.words(),
                    doc_author: faker.name.findName(),doc_status : faker.commerce.color(2),doc_fine:faker.commerce.price()},
                    {barcode : faker.phone.phoneNumberFormat(),id_doc:faker.random.number(60),doc_name: faker.random.words(),
                    doc_author: faker.name.findName(),doc_status : faker.commerce.color(2),doc_fine:faker.commerce.price()},
                    {barcode : faker.phone.phoneNumberFormat(),id_doc:faker.random.number(60),doc_name: faker.random.words(),
                    doc_author: faker.name.findName(),doc_status : faker.commerce.color(2),doc_fine:faker.commerce.price()}
               ],
               duration: faker.random.number(7),
               borrow_status: faker.random.number({
                'min': 0,
                'max': 2
            }),
            fee: faker.commerce.price()
           }
       }),
// phiếu mượn
       borrow_form: _.times(60, function(n){
        return{
            id:n,
            id_reader:faker.random.number(60),
            reader: faker.name.findName(),
            id_librarian: faker.random.number(60),
            docs: [
                {barcode : faker.phone.phoneNumberFormat()},
                {barcode : faker.phone.phoneNumberFormat()},
                {barcode : faker.phone.phoneNumberFormat()}
            ],
            duration: faker.random.number(7),
            borrow_status: faker.random.number({
             'min': 0,
             'max': 2
         })
        }
    }),
// độc giả vi phạm
    violated_user: _.times(60, function (n) {
        return {
          user_id: faker.random.number(60),
          user_name: faker.name.findName(),
          violate_time: n,
        };
      }),
    }       
}