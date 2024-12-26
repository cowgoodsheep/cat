<template>
    <el-row class="tac">
        <el-col :span="24">
            <el-menu active-text-color="#79bbff" background-color="#2b2c43" class="el-menu-vertical-demo"
                default-active="/main" text-color="#fff" @open="handleOpen" @close="handleClose" router="true">
                <h3>校园流浪猫管理</h3>
                <el-menu-item index="/home">
                    <el-icon>
                        <HomeFilled />
                    </el-icon>
                    <span>主页</span>
                </el-menu-item>
                <el-menu-item index="/catm">
                    <el-icon>
                        <View />
                    </el-icon>
                    <span>查看收录的猫猫</span>
                </el-menu-item>
                <el-menu-item v-if="role === 'admin'" index="/userm">
                    <el-icon>
                        <UserFilled />
                    </el-icon>
                    <span>查看用户列表</span>
                </el-menu-item>

                <el-menu-item index="/adaptm">
                    <el-icon>
                        <Lollipop />
                    </el-icon>
                    <span>查看领养系统</span>
                </el-menu-item>
                <el-menu-item v-if="role === 'user'" index="/adaptu">
                    <el-icon>
                        <Sugar />
                    </el-icon>
                    <span>申请领养系统</span>
                </el-menu-item>
            </el-menu>
        </el-col>
    </el-row>
</template>
  
<script lang="ts" setup>
import {
    Document,
    Menu as IconMenu,
    Location,
    Setting,
} from '@element-plus/icons-vue'
const role = localStorage.getItem("role")

const handleOpen = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
</script>

<style>
.tac {
    width: 200px;
    position: fixed;
}

el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 100vh;
}

.el-menu {
    height: 100vh;
    overflow: hidden;
    border: none;

    h3 {
        color: #fff;
        text-align: center;
        line-height: 48px;
        font-size: 16px;
        font-weight: 400;
    }
}
</style>
