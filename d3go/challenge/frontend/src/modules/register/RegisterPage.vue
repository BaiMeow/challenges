<template>
    <div class="register-container position-center">
        <div class="register-box">
            <h1 class="font-555">{{ $t('register.welcome') }}</h1>
            <el-form class="register-in" ref="registerForm" :model="registerForm" :rules="rules">
                <el-form-item prop="userName">
                    <label class="input-labels font-555">{{ $t('register.username') }}</label>
                    <el-input class="input" v-model="registerForm.userName" prefix-icon="User" clearable placeholder="Username"></el-input>
                </el-form-item>
                <el-form-item prop="passWord">
                    <label class="input-labels font-555">{{ $t('register.password') }}</label>
                    <el-input class="input" v-model="registerForm.passWord" prefix-icon="Bell" show-password clearable placeholder="Password"></el-input>
                </el-form-item>
                <el-form-item class="btn-box">
                    <el-button class="btn-register" type="primary" @click="submitForm('registerForm')">{{ $t('register.register') }}</el-button>
                </el-form-item>
            </el-form>
            <span class="signin font-555">{{ $t('register.hint') }}<router-link class="links" to="/loginPage">{{ $t('register.login') }}</router-link></span>
        </div>
    </div>
</template>


<script>
import axios from "axios"

export default {
    data() {
        return {
            registerForm: {
                userName: '',
                passWord: '',
            },
            rules: {
                userName: [
                    { required: true, message: this.$t('register.usernameNotice'), trigger: 'blur' },
                ],
                passWord: [
                    { required: true, message: this.$t('register.usernameNotice'), trigger: 'blur' },
                ]
            }
        }
    },
    methods: {
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    axios.post('/register', {
                            username: this.registerForm.userName,
                            password: this.registerForm.passWord
                        }, {
                            headers: {
                                'Content-Type': 'application/json'
                            }
                        })
                        .then(response => {
                            const data = response.data
                            if (data.status_code === 0) {
                                this.$message({
                                    type: 'success',
                                    message: data.status_msg
                                });
                                this.$router.push('/loginPage');
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
                        message: '参数不合法'
                    });

                    return false;
                }
            });
        }
    }
}
</script>


<style scoped>
.register-container {
    z-index: 2;
}

.register-box {
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

.btn-register {
    margin-top: 30px;
    height: 40px;
    width: 100%;
}

.signin {
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
