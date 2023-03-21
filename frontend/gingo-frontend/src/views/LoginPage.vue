<!-- eslint-disable no-import-assign -->
<!-- eslint-disable no-import-assign -->
<template>
    <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
        <el-form-item label="Email" prop="email">
            <el-input v-model="ruleForm.email" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="Password" prop="pass">
            <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')">Login</el-button>
            <el-button @click="resetForm('ruleForm')">Reset</el-button>
            <el-button @click="openRegisterPage()">Register</el-button>
        </el-form-item>
    </el-form>
</template>
 
<script>
import axios from "axios"
import { ElNotification } from 'element-plus'
import md5 from 'js-md5' 
    export default {
        data() {
            var validateEmail = (rule, value, callback) => {
                const mailReg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/;
                if (value === '') {
                    callback(new Error('请输入邮箱地址'));
                } else if(!mailReg.test(value)){
                    callback(new Error('请输入正确邮箱格式地址'));
                } else {
                    if (this.ruleForm.email !== '') {
                        this.$refs.ruleForm.validateField('email');
                    }
                    callback();
                }
            };
            var validatePass = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请输入密码'));
                } else {
                    if (this.ruleForm.pass !== '') {
                        this.$refs.ruleForm.validateField('pass');
                    }
                    callback();
                }
            };

            return {
                ruleForm: {
                    email:'',
                    pass: '',
                },
                rules: {
                    userName: [
                        { validator: validateEmail, trigger: 'blur' }
                    ],
                    pass: [
                        { validator: validatePass, trigger: 'blur' }
                    ]
                }
            };
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        var loginInfo = {};
                        loginInfo.Email = this.ruleForm.email;
                        loginInfo.Password = md5(this.ruleForm.pass)
                        axios.post('/api/login', loginInfo).then((resp) => {    
                            console.log(resp.data)
                            if (resp.data.Status == "failure") {
                                ElNotification({
                                    title: 'Login',
                                    message: resp.data.Message,
                                    position: 'bottom-right',
                                    duration: 0,
                                    type: 'error',
                                })
                            } else {
                                ElNotification({
                                    title: 'Login',
                                    message: resp.data.Message,
                                    position: 'bottom-right',
                                    type: 'success',
                                })

                                window.sessionStorage.setItem("currentUser", JSON.stringify(resp.data.Object))
                                this.openMainPage()
                            }
                            }).catch((err) => {
                                console.log(err)
                            })
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },

            resetForm(formName) {
                this.$refs[formName].resetFields();
            },

            openRegisterPage() {
                this.$router.push('/Register');
            },

            openMainPage() {
                this.$router.push('/Home');
            }
        }
    }
</script>
 
<style scoped>
</style> 
