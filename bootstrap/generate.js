module.exports = function(){
    var faker = require("faker")
    var _ = require("lodash")
    return {
       users: _.times(60, function(n){
           return {
            id: n,
            role: faker.lorem.word(),
            username: faker.internet.userName(),
            password: faker.internet.password(),
            name: faker.name.findName(),
            dob: faker.date.past(),
            sex: faker.lorem.word(),
            phone_number: faker.phone.phoneNumber(),
            created_at: faker.date.recent(),
            updated_at: faker.date.recent()
           }
       }),

    //    user_sec: _.times(60,function(n){
    //        return {
    //            id:n,
    //            username: faker.internet.userName(),
    //            password: faker.internet.password(),
    //            created_at: faker.date.recent(),
    //            updated_at: faker.date.recent()
    //        }
    //    }),
       
       documents: _.times(60, function(n){
        return {
            id:n,
            doc_name: faker.random.words(),
            doc_author: faker.name.findName(),
            doc_type: faker.commerce.color(),
            doc_description: faker.company.bs(),
            status: faker.lorem.word(),
            created_at: faker.date.recent(),
            updated_at: faker.date.recent(),
        }
       }),
       doc_status: _.times(60, function(n){
           return {
               id_borrow_form:n,
               status: faker.commerce.color(),
               created_at: faker.date.recent(),
               updated_at: faker.date.recent(),
               fee: faker.commerce.price()
           }
       }),
       borrow_form: _.times(60, function(n){
           return{
               id:n,
               id_customer:faker.random.number(),
               id_librarian: faker.random.number(60),
               id_doc: faker.random.number(60),
               duration: faker.random.number(7)
           }
       })
    }
}