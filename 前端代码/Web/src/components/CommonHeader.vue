<template>
    <div class="header-container">
        <div class="l-content">
            <span class="text" style="color:#606266;">
                Hi，{{ user_name }} ！ 欢迎来到校园流浪猫管理系统 ^v^
            </span>
        </div>
        <div class="r-content">
            <el-dropdown>
                <span class="el-dropdown-link">
                    <img class="user" :src="avatarUrl ? avatarUrl : localUserUrl" alt="">
                </span>
                <template #dropdown>
                    <el-dropdown-menu @click="goToFile">
                        <el-dropdown-item v-if="role === 'user'">个人中心</el-dropdown-item>
                    </el-dropdown-menu>
                    <el-dropdown-menu @click="logOut">
                        <el-dropdown-item>退出</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { useRouter } from 'vue-router';
import { ref, watch } from "vue"
import localUserUrl from "@/assets/images/user_default.jpg"
const $router = useRouter();
const user_name = ref<string | null>(localStorage.getItem("name"))
let avatarUrl = ref<string | null>(localStorage.getItem("avatarurl"))
const role = localStorage.getItem("role")
const goToFile = () => {
    $router.push('/personalfile');
}
const logOut = () => {
    //清除缓存
    window.localStorage.clear();
    $router.push('/login');
}
</script> 
<style scoped>
.header-container {
    background-color: #eceef2;
    width: calc(100vw - 240px);
    height: 65px;
    display: flex;
    z-index: 111;
    justify-content: space-between;
    align-items: center;
    padding: 0 20px;
    position: fixed;

    border-bottom: 5px solid #f9f9fa;

    .text {
        color: #242a2d;
        font-size: 14px;
        margin-left: 10px;
    }

    .r-content {

        .user {
            width: 50px;
            height: 50px;
            border-radius: 50%;
        }
    }
}
</style>
