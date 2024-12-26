<template>
    <div class="box" :model="tableData">
        <div class="cat-present">
            <el-card class="l-card" :body-style="{ padding: '0px' }">
                <img :src="tableData.cat_avatar" class="image" />
                <div style="padding: 60px">
                    <span v-if="tableData.cat_id != null">最新被收养的猫猫！</span>
                    <span v-else>还没有猫猫被收养！</span>
                    <div class="bottom">
                        <time class="time">{{ currentDate }}</time>
                        <el-button text class="button">Operating</el-button>
                    </div>
                </div>
            </el-card>
        </div>
        <div class="data-present">
            <el-card class="box-card">
                <template #header>
                    <div class="card-header">
                        <span style="font-size: 25px;color:#606266;">收录猫猫总数</span>
                    </div>
                </template>
                <div style="text-align: center;">
                    <span style="font-size: 25px;">现在一共有</span>
                    <span style="font-size: 35px;color:rgb(153, 140, 193);">{{ tableData.tot_amount }}</span>
                    <span style="font-size: 25px;">只猫猫被收录</span>
                </div>
            </el-card>
            <el-card class="box-card">
                <template #header>
                    <div class="card-header">
                        <span style="font-size: 25px;color:#606266;">猫猫男女比例</span>
                    </div>
                </template>
                <div style="text-align: center;">
                    <span style="font-size: 25px;">根据记录现在一共有:</span>
                    <span style="font-size: 35px;color:#dc9ca6;">{{ tableData.cat_female }}</span>
                    <span style="font-size: 25px;">只母猫,</span>
                    <span style="font-size: 35px;color:#8c9ec1;">{{ tableData.cat_male }}</span>
                    <span style="font-size: 25px;">只公猫</span>
                </div>
            </el-card>
            <el-card class="box-card">
                <template #header>
                    <div class="card-header">
                        <span style="font-size: 25px;color:#606266;">猫猫绝育总数</span>
                    </div>
                </template>
                <div style="text-align: center;">
                    <span style="font-size: 25px;">根据记录一共有</span>
                    <span style="font-size: 35px;color:#a3d392;">{{ tableData.cat_sterilization }}</span>
                    <span style="font-size: 25px;">只猫猫已绝育</span>
                </div>
            </el-card>
            <el-card class="box-card">
                <template #header>
                    <div class="card-header">
                        <span style="font-size: 25px;color:#666360;">猫猫被领养总数</span>
                    </div>
                </template>
                <div style="text-align: center;">
                    <span style="font-size: 25px;">根据记录一共有</span>
                    <span style="font-size: 35px;color:rgb(219, 152, 152);">{{ tableData.cat_adapt }}</span>
                    <span style="font-size: 25px;">只猫猫已被领养</span>
                </div>
            </el-card>

        </div>
    </div>
</template>
  
<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { get } from "@/api"
import catAvatarUrl from "@/assets/images/admin_default.jpg"

const currentDate = ref(new Date())
interface HomeInfo {
    "cat_id": number | null,
    "cat_name": null,
    "cat_avatar": string,
    "tot_amount": number | null,
    "cat_male": number | null,
    "cat_female": number | null,
    "cat_adapt": number | null,
    "cat_sterilization": number | null,
}
const tableData = ref<HomeInfo>({
    "cat_id": null,
    "cat_name": null,
    "cat_avatar": catAvatarUrl,
    "tot_amount": 0,
    "cat_male": 0,
    "cat_female": 0,
    "cat_adapt": 0,
    "cat_sterilization": 0,
});

onMounted(async () => {
    try {
        const result = await get('/home');
        console.log(result);
        tableData.value = result.data;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
});
</script>
  
<style>
.box {
    display: flex;
    justify-content: space-between;
    margin-top: 20px;
    border: 10px;
    /* width: calc(100vh - 250px); */

    .cat-present {
        display: flex;
        /* justify-content: center; */
        align-items: center;
        width: 40vw;

        .l-card {
            width: 100%;
            height: 83vh;
        }

        .time {
            font-size: 12px;
            color: #999;
        }

        .bottom {
            margin-top: 13px;
            line-height: 12px;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }

        .button {
            padding: 0;
            min-height: auto;
        }

        .image {
            width: 100%;
            display: block;
        }
    }

    .data-present {
        width: 40vw;
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
        align-content: space-between;

        .box-card {
            margin: 9px;
            width: 70vh;
            height: 19vh;
        }
    }
}



/*自带样式 */

.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.text {
    font-size: 14px;
}

.item {
    margin-bottom: 18px;
}

.box-card {
    width: 480px;
}
</style>
  