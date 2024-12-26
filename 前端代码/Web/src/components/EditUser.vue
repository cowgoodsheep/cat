<template>
    <el-form ref="ruleFormRef" :model="users" :rules="rules" label-width="120px" class="demo-users" :size="formSize"
        status-icon>
        <div label="头像" prop="user_avatarFile" class='upload-box'>
            <el-upload class="avatar-uploader" action="/api/uploaduser" :show-file-list="false"
                :before-upload="beforeAvatarUpload" :on-success="handleAvatarSuccess">
                <img v-if="imageUrl" :src="imageUrl" class="avatar" />
                <el-icon v-else class="avatar-uploader-icon">
                    <Plus />
                </el-icon>
            </el-upload>
        </div>
        <div class="u-title">上传头像</div>
        <el-form-item label="姓名" prop="user_name">
            <el-input v-model="users.user_name" type="input" placeholder="请输入用户名">
            </el-input>
        </el-form-item>
        <el-form-item label="密码" prop="user_password">
            <el-input v-model="users.user_password" type="password" placeholder="请输入密码">
            </el-input>
        </el-form-item>
        <el-form-item label="手机号码" prop="user_phone">
            <el-input v-model="users.user_phone" type="input" placeholder="请输入手机号码">
            </el-input>
        </el-form-item>
        <el-form-item label="年龄" prop="user_age">
            <el-input-number v-model="users.user_age" :min="0" :max="130">
            </el-input-number>
        </el-form-item>
        <el-form-item label="性别" prop="user_gender">
            <el-radio-group v-model="users.user_gender">
                <el-radio label="男" />
                <el-radio label="女" />
            </el-radio-group>
        </el-form-item>

        <el-form-item>
            <el-button type="primary" @click="submitForm(ruleFormRef)">
                确认
            </el-button>
        </el-form-item>
    </el-form>
</template>
  
<script lang="ts" setup>
import { reactive, ref, defineProps, onMounted, watch } from 'vue'
import type { FormInstance, FormRules, UploadProps } from 'element-plus'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { get, post, put } from '@/api';
import { AxiosHeaders } from 'axios'

const imageUrl = ref<string>('');
const emit = defineEmits();
interface UserInfo {
    user_name: string,
    user_password: string,
    user_phone: string,
    user_age: number | null,
    user_gender: string,
    user_avatar: string,
}

const formSize = ref('default')
const ruleFormRef = ref<FormInstance>()
const props = defineProps({
    //子组件接收父组件传递过来的值
    editItem: Object,
})

onMounted(() => {
    console.log(props)
})
const users = reactive<UserInfo>({
    user_name: '',
    user_password: '',
    user_phone: '',
    user_age: null,
    user_gender: '',
    user_avatar: '',
})
watch(() => props?.editItem, (newVal) => {
    console.log(props?.editItem);
    Object.assign(users, newVal)
}, { deep: true, immediate: true })

const rules = reactive<FormRules<UserInfo>>({
    user_name: [
        { required: true, message: '请输入用户的名字', trigger: 'blur' },
        { min: 0, max: 20, message: '最多20字', trigger: 'blur' },
    ],
    user_password: [
        { required: true, message: '请输入用户的密码', trigger: 'blur' },
        { min: 6, max: 20, message: '长度为6~20位', trigger: 'blur' },
    ],
    user_phone: [
        { required: false, message: '请输入用户的手机号码', trigger: 'blur' },
        { min: 11, max: 11, message: '长度为11位', trigger: 'blur' },
    ],
    user_gender: [
        {
            required: false,
            message: '请选择用户的性别',
            trigger: 'change',
        },
    ],
    user_age: [
        {
            required: true,
            message: '请输入您的年龄',
            trigger: 'change',
        },
    ],

})
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
    users.user_avatar = response.data.url
}

const submitForm = async (formEl: FormInstance | undefined) => {
    console.log(users, "tetetetet");
    if (!formEl) return
    await formEl.validate(async (valid, fields) => {
        if (valid) {
            console.log('submit!')
            try {
                const result = await put('/userm', users);
                console.log(result);
                ElMessage.success("修改成功！")
            } catch (error) {
                console.error('Error fetching data:', error);
                ElMessage.error("修改失败！")
            }
            emit("userEdited")
        } else {
            console.log('error submit!', fields)
        }
    })
    formEl.resetFields()
    imageUrl.value = ''
}

</script>

<style scoped>
.upload-box {
    display: flex;
    justify-content: center;
    align-content: space-between;
    margin: 15px;
}

.u-title {
    display: flex;
    justify-content: center;
    margin: 15px;
}

.avatar-uploader .avatar {
    width: 178px;
    height: 178px;
    display: block;
}
</style>

<style>
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