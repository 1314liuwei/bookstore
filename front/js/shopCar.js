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
        handleShopping(){
            
            console.log("加入购物车")
            if(!flag || !username){
                alert("请登录后再使用")
                return
            }

            console.log("cookie: ",document.cookie)
            console.log("cookie")

            axios.defaults.withCredentials = true
            
            axios.post("http://116.62.229.23:8080/shopping_cart/addBook").then(res => {
                console.log(res)
            })

        }
    }
})



var filter = new Vue({
    el:"#filter",
    data:{
        categoriesList:[],

    },
    methods:{
        getFilter(){
            axios.defaults.withCredentials = true
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

selection.getCategory()

filter.getFilter()
filter.handleCate()