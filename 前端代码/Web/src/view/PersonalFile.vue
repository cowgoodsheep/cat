<template>
    <div v-if="!isEdit" style="display: flex;justify-content:center" :model="userData">
        <el-card style="display:flex;height:88vh;width:75vw;flex-flow:column;">
            <div style="margin:20px;padding:20px;">
                <div>
                    <h2 style="text-align:center;">欢迎来到个人中心</h2>
                </div>
            </div>
            <div style="margin:20px;padding:50px 20px 20px 20px;display:flex;">
                <div style="margin:auto 50px;">
                    <el-avatar :size="280" :src="userData.user_avatar" alt='' />
                </div>
                <div style="margin-left: 11vw;width:45vw;display:flex;flex-flow:column;">
                    <el-card style="margin:9px;">
                        <div style="display: flex;">
                            <User style="width: 1.5em;" />
                            <span style="margin-left: 1em;">用户名:</span>
                            <span style="margin-left: auto;">{{ userData.user_name }}</span>
                        </div>
                    </el-card>
                    <el-card style="margin:9px;">
                        <div style="display: flex;">
                            <Coordinate style="width: 1.5em;" />
                            <span style="margin-left: 1em;">性别:</span>
                            <span style="margin-left: auto;">{{ userData.user_gender }}</span>
                        </div>
                    </el-card>
                    <el-card style="margin:9px;">
                        <div style="display: flex;">
                            <Calendar style="width: 1.5em;" />
                            <span style="margin-left: 1em;">年龄:</span>
                            <span style="margin-left: auto;">{{ userData.user_age }}</span>
                        </div>
                    </el-card>
                    <el-card style="margin:9px;">
                        <div style="display: flex;">
                            <Cellphone style="width: 1.5em;" />
                            <span style="margin-left: 1em;">手机号码:</span>
                            <span style="margin-left: auto;">{{ userData.user_phone }}</span>
                        </div>
                    </el-card>
                    <el-card style="margin:9px;">
                        <div style="display: flex;">
                            <Sugar style="width: 1.5em;" />
                            <span style="margin-left: 1em;">领养的猫猫编号:</span>
                            <span style="margin-left: auto;">{{ userData.user_cat }}</span>
                        </div>
                    </el-card>
                </div>
            </div>
            <div style="text-align:right;margin-right:60px;">
                <el-button type="primary" @click="Editinit">
                    修改
                </el-button>
            </div>
        </el-card>
    </div>
    <div v-else style="display: flex;justify-content:center">
        <el-card style="display:flex;height:88vh;width:75vw;flex-flow:column;">
            <div label="头像" prop="user_avatarFile" class='upload-box'>
                <el-upload class="avatar-uploader" action="/api/uploaduser" :show-file-list="false"
                    :before-upload="beforeAvatarUpload" :on-success="handleAvatarSuccess">
                    <img v-if="imageUrl" :src="imageUrl" class="avatar" />
                    <el-icon v-else class="avatar-uploader-icon">
                        <Plus />
                    </el-icon>
                </el-upload>
            </div>
            <el-form ref="ruleFormRef" :model="editData" :rules="rules" label-width="120px" class="demo-rulForm"
                :size="formSize" status-icon style="margin-left: 25%;">
                <el-form-item label="姓名" prop="user_name">
                    <el-input v-model="editData.user_name" />
                </el-form-item>
                <el-form-item label="密码" prop="user_password">
                    <el-input v-model="editData.user_password" />
                </el-form-item>
                <el-form-item label="性别" prop="user_gender">
                    <el-radio-group v-model="editData.user_gender">
                        <el-radio label="男" />
                        <el-radio label="女" />
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="年龄" prop="user_age">
                    <el-input v-model="editData.user_age" />
                </el-form-item>
                <el-form-item label="手机号码" prop="user_phone">
                    <el-input v-model="editData.user_phone" />
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="submitForm(ruleFormRef)">
                        提交
                    </el-button>
                    <el-button type="info" @click="clickReturn">
                        返回
                    </el-button>
                </el-form-item>
            </el-form>
        </el-card>
    </div>
</template>
  
<script lang="ts" setup>
import { reactive, ref, onMounted } from 'vue'
import type { FormInstance, FormRules, UploadProps } from 'element-plus'
import pho from "@/assets/images/user.jpg"
import UploadAvatar from "@/components/UploadAvatar.vue";
import { get, put } from "@/api"
import { ElMessage } from "element-plus"
import { useRouter } from "vue-router"
import { AxiosError } from "axios"
const router = useRouter()
interface RuleForm {
    name: string,
    gender: string,
    age: string,
    phone: string,
    password: string,
}

const formSize = ref('default')
const ruleFormRef = ref<FormInstance>()
const idFromLocalStorage: string | null = localStorage.getItem("id");
const id = idFromLocalStorage ? parseInt(idFromLocalStorage, 10) : 0;
const ruleForm = reactive<RuleForm>({
    name: '',
    gender: '',
    age: '',
    phone: '',
    password: '',
})
interface UserInfo {
    user_id: number | null,
    user_name: string,
    user_password: string,
    user_phone: string,
    user_age: string,
    user_gender: string,
    user_avatar: string,
    user_cat: string,
}
const userData = ref<UserInfo>({
    user_id: null,
    user_name: '',
    user_password: '',
    user_phone: '',
    user_age: '',
    user_gender: '',
    user_avatar: '',
    user_cat: '',
});
const editData = ref<UserInfo>({
    user_id: id,
    user_name: '',
    user_password: '',
    user_phone: '',
    user_age: '',
    user_gender: '',
    user_avatar: '',
    user_cat: '',
});
const rules = reactive<FormRules<RuleForm>>({
    name: [
        { required: true, message: '请输入您的姓名', trigger: 'blur' },
        { min: 0, max: 20, message: '最多20字', trigger: 'blur' },
    ],
    gender: [
        {
            required: true,
            message: '请选择您的性别',
            trigger: 'change',
        },
    ],
    age: [
        {
            required: true,
            message: '请输入您的年龄',
            trigger: 'change',
        },
    ],


})

const isEdit = ref(false)
const imageUrl = ref<string>('');
const responseUrl = ref<string>('');

const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
    if (rawFile.type !== 'image/jpeg') {
        ElMessage.error('Avatar picture must be JPG format!')
        return false
    } else if (rawFile.size / 1024 / 1024 > 10) {
        ElMessage.error('Avatar picture size can not exceed 10MB!')
        return false
    }
    return true
}
const handleAvatarSuccess: UploadProps['onSuccess'] = (
    response,
    uploadFile
) => {
    imageUrl.value = URL.createObjectURL(uploadFile.raw!)
    responseUrl.value = response.data.url
}
const clickReturn = async () => {
    imageUrl.value = ''
    isEdit.value = false
    try {
        const url = '/personalfile?id=' + idFromLocalStorage
        const result = await get(url);
        userData.value = result.data
    } catch (error) {
        console.error('Error fetching data:', error);
    }
}
const Editinit = () => {
    isEdit.value = true
    editData.value = userData.value
}
const submitForm = async (formEl: FormInstance | undefined) => {
    if (!formEl) return
    await formEl.validate(async (valid, fields) => {
        if (valid) {
            try {
                editData.value.user_id = id
                editData.value.user_avatar = responseUrl.value
                const result = await put("/personalfile", editData.value)
                console.log(result.data)
                ElMessage.success("修改成功！")
                localStorage.setItem('name', editData.value.user_name)
                if (editData.value.user_avatar != '') {
                    localStorage.setItem('avatarurl', editData.value.user_avatar)
                }
                isEdit.value = false;
                imageUrl.value = ''
                router.go(0);
            } catch (error) {
                // 使用类型断言将 error 断言为 AxiosError 类型
                const axiosError = error as AxiosError<{ data: { msg: string } }>;
                console.log(axiosError);
                // 如果是 AxiosError，并且包含响应信息
                if (axiosError.isAxiosError && axiosError.response) {
                    const errorMsg = axiosError.response.data.data.msg;
                    ElMessage.error(`修改失败：${errorMsg}`);
                    // 显示具体的后端错误信息
                } else {
                    // 显示通用的错误信息
                    ElMessage.error('修改失败：未知错误！');
                }
            }
            console.log('submit!')
        } else {
            console.log('error submit!', fields)
            ElMessage.error("修改失败！")
        }
    })
}

onMounted(async () => {
    try {
        const url = '/personalfile?id=' + idFromLocalStorage
        const result = await get(url);
        console.log(result)
        userData.value = result.data
        console.log(userData, 111)
    } catch (error) {
        console.error('Error fetching data:', error);
    }
})

</script>
<style scoped>
/* .el-descriptions-item {
    width: 10vw;
} */

.text-box {
    font-size: 20px !important;
}

.el-descriptions-item {
    max-width: 200px;
}

.el-form-item {
    width: 500px;
}

.avatar-uploader .avatar {
    width: 178px;
    height: 178px;
    display: block;
}
</style>

<style>
.upload-box {
    display: flex;
    justify-content: center;
    align-content: space-between;
    margin: 15px;
}

.avatar-uploader .el-upload {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
    border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
}
</style>