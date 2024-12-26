<template>
    <div class="common-layout">
        <el-container calss>
            <el-aside width="200px">
                <common-aside />
            </el-aside>
            <el-container>
                <el-header>
                    <common-header />
                </el-header>
                <el-main class="main-box">
                    <router-view />
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>
  
<style>
.main-box {
    background-color: #eceef2;
    height: calc(100vh - 60px);
}

.el-header {
    padding: 0;
}
</style>
<script setup lang="ts">
import CommonAside from '../components/CommonAside.vue'
import CommonHeader from '../components/CommonHeader.vue'
</script>
