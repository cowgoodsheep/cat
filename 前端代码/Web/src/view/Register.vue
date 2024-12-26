<template>
    <el-form :model="form" class="login-container">
        <h3>系统注册</h3>
        <el-form-item label="姓名:">
            <el-input v-model="form.user_name" type="input" placeholder="请输入用户名">
            </el-input>
        </el-form-item>
        <el-form-item label="密码:">
            <el-input v-model="form.user_password" type="password" placeholder="请输入密码">
            </el-input>
        </el-form-item>
        <el-form-item label="手机号码:">
            <el-input v-model="form.user_phone" type="input" placeholder="请输入手机号码">
            </el-input>
        </el-form-item>
        <el-form-item label="年龄:">
            <el-input-number v-model="form.user_age" :min="0" :max="130">
            </el-input-number>
        </el-form-item>
        <el-form-item label="性别:">
            <el-radio-group v-model="form.user_gender">
                <el-radio label="男" />
                <el-radio label="女" />
            </el-radio-group>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="sendData">提交</el-button>
            <el-button @click="goToLogin">返回</el-button>
        </el-form-item>
    </el-form>
</template>
<script lang = "ts" setup>
import { reactive } from 'vue'
import { useRouter } from 'vue-router';
import { get, post } from '@/api';
import { AxiosError } from "axios"
import { ElMessage } from "element-plus"
const $router = useRouter();
interface Form {
    user_name: string
    user_password: string
    user_phone: string
    user_age: number
    user_gender: string
}
const form = reactive<Form>({} as Form)

async function sendData() {
    try {
        const result = await post('/register', form);
        console.log(result);
        ElMessage.success(result.data.msg)
    } catch (error) {
        console.log(form);
        // 使用类型断言将 error 断言为 AxiosError 类型
        const axiosError = error as AxiosError<{ data: { msg: string } }>;
        console.log(axiosError);

        // 如果是 AxiosError，并且包含响应信息
        if (axiosError.isAxiosError && axiosError.response) {
            const errorMsg = axiosError.response.data.data.msg;
            ElMessage.error(`注册失败：${errorMsg}`);
            // 显示具体的后端错误信息
        } else {
            // 显示通用的错误信息
            ElMessage.error('注册失败，请检查你的信息');
        }
    }
}
const goToLogin = () => {
    $router.push('/login');
}
</script>
<style scoped>
.login-container {
    width: 350px;
    background-color: #fff;
    border: 1px solid#eaeaea;
    border-radius: 15px;
    padding: 35px 35px 15px 35px;
    box-shadow: 0 0 25px #cacaca;
    margin: 180px auto;

    h3 {
        text-align: center;
        margin-bottom: 20px;
        color: #505450;
    }

    /deep/.el-form-item__content {
        justify-content: center;
    }

}
</style>