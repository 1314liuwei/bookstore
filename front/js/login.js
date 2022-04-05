var obj = {
    data() {
        return {
            name: '',
            pwd: '',
        }
    },
    methods: {
        handleLogin() {
            console.log('yes')
            console.log(this.name, this.pwd)
            axios.defaults.withCredentials = true
            axios.post("http://116.62.229.23:8080/user/sign_in", {
                username: this.name,
                password: this.pwd
            }).then(res => {
                axios.defaults.withCredentials = true
                console.log(document.cookie)
                if (res.data.code === 0) {
                    // console.log(res)
                    // window.localStorage.setItem('flag', true)
                    // window.localStorage.setItem('name',this.name)
                    // location.href = './index.html'
                } else {
                    alert("请检查用户名和密码是否输入正确")
                }
            })

        }
    }
}
Vue.createApp(obj).mount('#login')