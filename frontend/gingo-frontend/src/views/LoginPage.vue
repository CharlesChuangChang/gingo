<template>
    <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
        <el-form-item label="UserName/Email" prop="userName">
            <el-input v-model="ruleForm.userName" autocomplete="off"></el-input>
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
    export default {
        data() {
            var validateUserName = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请输入用户名或者邮箱地址'));
                } else {
                    if (this.ruleForm.userName !== '') {
                        this.$refs.ruleForm.validateField('userName');
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
                    userName:'',
                    pass: '',
                },
                rules: {
                    userName: [
                        { validator: validateUserName, trigger: 'blur' }
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
                        this.$message({
                            message: '提交成功',
                            type:'success'
                        });
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
            }
        }
    }
</script>
 
<style scoped>
</style> 
