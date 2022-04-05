var obj={
    data(){
        return{
            username:'',
            password:'',
            confirm:'',
            message:'',
        }
    },
    methods:{
        handleRegister(){
            axios.post("http://116.62.229.23:8080/user/sign_up",{
                username:this.username,
                password:this.password,
                password2:this.confirm
            }).then(res => {
                console.log(res)
                if(res.data.code===0){
                    alert("注册成功")
                    window.location.href = './login.html'
                }else{
                    alert(res.data.message+`\n请重新输入账号密码`)

                }
            })
        }
    }

}
Vue.createApp(obj).mount('#register')