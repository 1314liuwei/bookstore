// const { Axios, default: axios } = require("axios");


const flag = window.localStorage.getItem('flag')
const username = window.localStorage.getItem('name')

const expire = 1000 * 60*30;
setTimeout(() => {
    window.localStorage.setItem('flag', false)
    window.localStorage.setItem('flag', '')
    console.log("this is")
}, expire)

var vm = new Vue({
    el: "#index",
    data: {
        usersname: username,
        classobjoff: {
            active: true,
        },
        classobjon: {
            active: false
        }
    },

})

if (!flag) {
    console.log("失败")
} else {
    console.log("成功")
    console.log(vm.classobjoff.active)
    vm.classobjoff.active = false
    console.log(vm.classobjoff.active)
    vm.classobjon.active = true
}


var selection = new Vue({
    el: "#goods",
    data: {
        goodsList:[],
    },
    methods: {
        getCategory(){
            axios.post("http://116.62.229.23:8080/book/category").then(res => {
                this.goodsList=res.data.data.Books
                // console.log(res.data.data.Books)
            })
        },
        handleDetail(item){
            console.log("详情")
            console.log(item.id)
            window.localStorage.setItem('goodsId',item.id)
            window.location.href="./detail.html"
        },
        handleShopping(item){
            if(!flag || !username){
                alert("请登录后再使用")
                return
            }
            console.log("加入购物车")
            console.log(item.id)

            axios.defaults.withCredentials = true

            axios.defaults.withCredentials = true
            axios.post("http://116.62.229.23:8080/shopping_cart/addBook", {
                books:[item.id]
            }).then(res => {
                console.log(res)
            })
            
        }
    }
})
selection.getCategory()


var filter = new Vue({
    el:"#filter",
    data:{
        categoriesList:[],

    },
    methods:{
        getFilter(){
            axios.get("http://116.62.229.23:8080/category/all").then(res => {
                // console.log(res.data.data.categories)
                this.categoriesList=res.data.data.categories
            })
        },
        handleCate(item){
            axios.defaults.withCredentials = true
            axios.post("http://116.62.229.23:8080/book/category",{
                category:item
            }).then(res => {
                console.log(res.data.data.Books)
                selection.goodsList=res.data.data.Books
            })
        }
    }
})

filter.getFilter()
filter.handleCate()
