
var info = new Vue({
    el:"#content",
    data:{
        goodsId:window.localStorage.getItem('goodsId'),
        bookname:'',
        bookimg:'',
        bookprice:'',
    },
    methods:{
        getInfo(){
            console.log("成功")
            axios.post("http://116.62.229.23:8080/book/id",{
                id:this.goodsId,
            }).then(res => { 
                console.log(res.data.data.Books)
                this.bookimg=res.data.data.Books[0].cover
                this.bookname=res.data.data.Books[0].name
                this.bookprice=res.data.data.Books[0].price
                console.log(this.bookimg, this.bookname, this.bookprice)
            })
        }
    }
})

info.getInfo()
