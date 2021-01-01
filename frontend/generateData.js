module.exports = function(){
    var faker = require("faker")
    var _ = require("lodash")
    
    return {
       documents: _.times(60, function(n){
        return {
            id:n,
            barcode : faker.phone.phoneNumberFormat(),
            doc_name: faker.random.words(),
            doc_author: faker.name.findName(),
            doc_type: faker.commerce.color(),
            doc_publisher: faker.company.bs(10),
            doc_edition : faker.random.number(),
            doc_price : faker.commerce.price(),
            doc_status : faker.commerce.color(2),
            created_at: faker.date.recent(),
            updated_at: faker.date.recent(),
        }
       }),
       doc_damaged: _.times(60, function(n){
           return {
               id_doc:n,
               barcode : faker.phone.phoneNumberFormat(),
               doc_name: faker.random.words(),
               doc_author: faker.name.findName(),
               librarian:faker.name.findName(),
               date: faker.date.recent(),
               fee: faker.commerce.price()
           }
       }),
       borrow_form: _.times(60, function(n){
           return{
               id:n,
               id_reader:faker.random.number(60),
               reader: faker.name.findName(),
               id_librarian: faker.random.number(60),
               librarian:faker.name.findName(),
               id_doc: faker.random.number(60),
               docs: [
                   {barcode : faker.phone.phoneNumberFormat(),id_doc:faker.random.number(60),doc_name: faker.random.words(),
                    doc_author: faker.name.findName()},
                    {barcode : faker.phone.phoneNumberFormat(),id_doc:faker.random.number(60),doc_name: faker.random.words(),
                    doc_author: faker.name.findName()},
                    {barcode : faker.phone.phoneNumberFormat(),id_doc:faker.random.number(60),doc_name: faker.random.words(),
                    doc_author: faker.name.findName()}
               ],
               duration: faker.random.number(7)
           }
       })
    }
}