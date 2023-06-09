import { createSlice } from '@reduxjs/toolkit'
import axios from "axios";
import {sharedSlice} from "./shared";

export const applicationSlice = createSlice({
    name: 'applicationSlice',
    initialState: {
        axios: null,
        user: {
            email: null,
            authorized: false
        }
    },
    reducers: {
        initAxios: state => {
            let instance = axios.create({
                baseURL: process.env.REACT_APP_AXIOS_BASE_URL
            })
            state.axios = instance
        },
        setToken: (state, data) => {
            state.axios.defaults.headers.common['Authorization'] = `Bearer ${data.payload.token}`;

            //set user
            state.user = {
                email: data.payload.email,
                authorized: true,
            }

            //set to sessions storage
            sessionStorage.setItem("account", JSON.stringify(data.payload))
        }
    }
})

export const { initAxios, setToken } = applicationSlice.actions

export default applicationSlice.reducer