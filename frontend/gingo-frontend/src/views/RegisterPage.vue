<template>
    <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
        <el-form-item label="UserName" prop="userName">
            <el-input v-model="ruleForm.userName" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="Email" prop="email">
            <el-input v-model="ruleForm.email" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="Password" prop="pass">
            <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="Confirm" prop="passConfirm">
            <el-input type="password" v-model="ruleForm.passConfirm" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item>
            <el-button type="primary" @click="submitForm()">Register</el-button>
            <el-button @click="resetForm()">Reset</el-button>
            <el-button @click="openLoginPage()">Login</el-button>
        </el-form-item>
    </el-form>
</template>
 
<script>
import axios from "axios"
import { ElNotification } from 'element-plus'
import md5 from 'js-md5' 
    export default {
        data() {
            var checkUserName = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请输入用户名'));
                } else {
                    if (this.ruleForm.userName !== '') {
                        this.$refs.ruleForm.validateField('userName');
                    }
                    callback();
                }
            };
            var checkEmail = (rule, value, callback) => {
                const mailReg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/;
                if (value === '') {
                    callback(new Error('请输入邮箱地址'));
                } else if (!mailReg.test(value)) {
                    callback(new Error("请输入正确的邮箱格式"));
                } else {
                    if (this.ruleForm.userEmail !== '') {
                        this.$refs.ruleForm.validateField('email');
                    }
                    callback();
                }
            };

            var validatePass = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请输入密码'));
                } else {
                    if (this.ruleForm.checkPass !== '') {
                        this.$refs.ruleForm.validateField('pass');
                    }
                    callback();
                }
            };
            var validatePass2 = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请再次输入密码'));
                } else if (value !== this.ruleForm.pass) {
                    callback(new Error('两次输入密码不一致!'));
                } else {
                    callback();
                }
            };
            return {
                ruleForm: {
                    userName:'',
                    email: '',
                    pass: '',
                    passConfirm: '',
                    age: ''
                },
                rules: {
                    userName: [
                        { validator: checkUserName, trigger: 'blur' }
                    ],
                    email: [
                        { validator: checkEmail, trigger: 'blur' }
                    ],
                    pass: [
                        { validator: validatePass, trigger: 'blur' }
                    ],
                    passConfirm: [
                        { validator: validatePass2, trigger: 'blur' }
                    ]
                }
            };
        },
        methods: {
            submitForm() {
                this.$refs.ruleForm.validate((valid) => {
                    if (valid) {
                        var registerInfo = JSON.parse(JSON.stringify(this.ruleForm))
                        registerInfo.pass  = md5(registerInfo.pass)
                        registerInfo.passConfirm  = md5(registerInfo.passConfirm)
                        axios.post('/api/register', registerInfo).then((resp) => {    
                            console.log(resp.data)
                            if (resp.data.Status == "failure") {
                                ElNotification({
                                    title: 'Register',
                                    message: resp.data.Message,
                                    position: 'bottom-right',
                                    duration: 0,
                                    type: 'error',
                                })
                            } else {
                                ElNotification({
                                    title: 'Register',
                                    message: resp.data.Message,
                                    position: 'bottom-right',
                                    type: 'success',
                                })
                                this.openLoginPage()
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
            resetForm() {
                this.$refs.ruleForm.resetFields();
            },
            openLoginPage() {
                this.$router.push('/Login');
            }
        }
    }
</script>
 
<style scoped>
</style> 
