<template>
    <div class="login-container position-center">
        <div class="login-box">
            <h1 class="font-555">{{ $t('login.welcome') }}</h1>
            <el-form class="login-in" ref="loginForm" :model="loginForm" :rules="rules">
                <el-form-item prop="userName">
                    <label class="input-labels font-555">{{ $t('login.username') }}</label>
                    <el-input class="input" v-model="loginForm.userName" prefix-icon="User" clearable placeholder="Username"></el-input>
                </el-form-item>
                <el-form-item prop="passWord">
                    <label class="input-labels font-555">{{ $t('login.password') }}</label>
                    <el-input class="input" v-model="loginForm.passWord" prefix-icon="Bell" show-password clearable placeholder="Password"></el-input>
                </el-form-item>
                <el-form-item class="btn-box">
                    <el-button class="btn-login" type="primary" @click="submitForm('loginForm')">{{ $t('login.login') }}</el-button>
                </el-form-item>
            </el-form>
            <span class="signup font-555">{{ $t('login.hint') }}<router-link class="links" to="/registerPage">{{ $t('login.register') }}</router-link></span>
        </div>
    </div>
</template>


<script>
import axios from "axios"

export default {
    data() {
        return {
            loginForm: {
                userName: '',
                passWord: '',
            },
            rules: {
                userName: [
                    { required: true, message: this.$t('login.usernameNotice'), trigger: 'blur' },
                ],
                passWord: [
                    { required: true, message: this.$t('login.passwordNotice'), trigger: 'blur' },
                ]
            }
        }
    },
    methods: {
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    axios.post('/login', {
                            username: this.loginForm.userName,
                            password: this.loginForm.passWord
                        }, {
                            headers: {
                                'Content-Type': 'application/json'
                            }
                        })
                        .then(response => {
                            const data = response.data
                            if (data.status_code === 0) {
                                this.$router.push('/uploadPage');
                            } else {
                                this.$message({
                                    type: 'error',
                                    message: data.status_msg
                                });
                            }
                        })
                        .catch(error => {
                            console.log(error);
                        });
                } else {
                    this.$message({
                        type: 'error',
                        message: this.$t('login.panic')
                    });

                    return false;
                }
            });
        }
    }
}
</script>


<style scoped>
.login-container {
    z-index: 2;
}

.login-box {
    width: 400px;
    height: fit-content;
    padding: 30px 40px;
    background-color: #fff;
    border-radius: 5px;
    box-shadow: 0 0 10px #ccc;
    box-sizing: border-box;
}

.font-555 {
    color: #555;
}

h1 {
    margin: 0 0 20px 0;
    padding: 0;
    font-size: 22px;
    font-weight: 400;
}

.input-labels {
    margin-bottom: 4px;
}

.input {
    height: 40px;
}

.btn-login {
    margin-top: 30px;
    height: 40px;
    width: 100%;
}

.signup {
    position: absolute;
    left: 20px;
    bottom: -40px;
    font-size: 14px;
}

.links {
    padding-left: 10px;
    text-decoration: none;
    color: #469bf0;
}
</style>
