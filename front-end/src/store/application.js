import { createSlice } from '@reduxjs/toolkit'
import axios from "axios";
import {sharedSlice} from "./shared";

export const applicationSlice = createSlice({
    name: 'applicationSlice',
    initialState: {
        axios: axios.create({
            baseURL: process.env.REACT_APP_AXIOS_BASE_URL
        }),
        user: {
            email: null,
            authorized: false
        },
        sideClosed: true,
        darkMode: false
    },
    reducers: {
        setToken: (state, data) => {
            state.axios.defaults.headers.common['Authorization'] = `Bearer ${data.payload.token}`;

            //set user
            state.user = {
                email: data.payload.email,
                authorized: true,
            }

            //set to sessions storage
            sessionStorage.setItem("account", JSON.stringify(data.payload))
        },
        logout: state => {
            //remove header
            delete state.axios.defaults.headers.common["Authorization"]

            //set user
            state.user = {
                email: null,
                authorized: false,
            }

            //Remove from sesion storage
            sessionStorage.removeItem("account")
        },
        setSideClosed: (state,data) => {
            if (data.payload != undefined) {
                state.sideClosed = data.payload
                return
            }
            state.sideClosed = !state.sideClosed
        },
        setDarkMode: (state, data) => {
            state.darkMode = data.payload
            sessionStorage.setItem("darkMode", data.payload)
        }
    }
})

export const {setToken, setSideClosed, logout, setDarkMode} = applicationSlice.actions

export default applicationSlice.reducer