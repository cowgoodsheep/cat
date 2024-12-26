<template>
    <div class="login-box">
        <el-form :model="form" class="login-container">
            <h3>校园流浪猫管理系统</h3>
            <div>
                <el-form-item size="large">
                    <el-input v-model="form.user_name" type="input" placeholder="请输入用户名">
                    </el-input>
                </el-form-item>
                <el-form-item size="large">
                    <el-input v-model="form.user_password" type="password" placeholder="请输入密码">
                    </el-input>
                </el-form-item>
            </div>
            <div>
                <el-form-item>
                    <el-button type="primary" @click="sendData">登录</el-button>
                    <el-button @click="goToRegister">注册</el-button>
                </el-form-item>
            </div>
        </el-form>
    </div>
</template>
<script lang = "ts" setup>
import { h, reactive } from 'vue'
import { post } from '@/api';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus'
import { AxiosError } from 'axios';
import UserAvatarDefault from "@/assets/images/user_default.jpg"
import AdminAvatarDefault from "@/assets/images/admin_default.jpg"
const $router = useRouter();

const form = reactive({
    user_name: '',
    user_password: '',
})
async function sendData() {
    try {
        const result = await post('/login', form);
        console.log(result);
        ElMessage.success(`欢迎您！${result.data.name}`)
        // 存数据到localStorage,然后全局都可以取用
        localStorage.setItem('role', result.data.role);
        // const role = localStorage.getItem('role');
        // console.log(role);
        localStorage.setItem('id', result.data.id)
        localStorage.setItem('name', result.data.name)
        if (result.data.avatar) {
            localStorage.setItem('avatarurl', result.data.avatar)
        } else if (result.data.role == "admin") {
            localStorage.setItem('avatarurl', AdminAvatarDefault)
        } else {
            localStorage.setItem('avatarurl', UserAvatarDefault)
        }

        $router.push('/home');
    } catch (error) {
        // 使用类型断言将 error 断言为 AxiosError 类型
        const axiosError = error as AxiosError<{ data: { msg: string } }>;
        console.log(axiosError);
        // 如果是 AxiosError，并且包含响应信息
        if (axiosError.isAxiosError && axiosError.response) {
            const errorMsg = axiosError.response.data.data.msg;
            ElMessage.error(`登录失败：${errorMsg}`);
            // 显示具体的后端错误信息
        } else {
            // 显示通用的错误信息
            ElMessage.error('登录失败，请检查你的用户名或者密码');
        }
    }
}
const goToRegister = () => {
    $router.push('/register');
}

</script>
<style scoped>
.login-box {
    position: relative;
    height: 100vh;
    background-image: url('@/assets/images/background.jpg');
}

.login-container {
    position: absolute;
    top: 4.5vw;
    right: 11vw;
    width: 350px;
    height: 275px;
    background-color: rgb(255, 255, 255, 0.5);
    border: 1px solid#eaeaea;
    border-radius: 15px;
    padding: 35px 35px 15px 35px;
    box-shadow: 0 0 25px #cacaca;
    margin: 180px auto;
    display: flex;
    flex-flow: column;
    justify-content: space-around;

    h3 {
        text-align: center;
        /* margin-bottom: 20px; */
        color: #666
    }

    /deep/.el-form-item__content {
        justify-content: center;
    }
}
</style>
