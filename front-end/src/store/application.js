import { createSlice } from '@reduxjs/toolkit'
import axios from "axios";
import {sharedSlice} from "./shared";

export const applicationSlice = createSlice({
    name: 'applicationSlice',
    initialState: {
        axios: null
    },
    reducers: {
        initAxios: state => {
            let instance = axios.create({
                baseURL: process.env.REACT_APP_AXIOS_BASE_URL
            })
            state.axios = instance
        }
    }
})

export const { initAxios } = applicationSlice.actions

export default applicationSlice.reducer