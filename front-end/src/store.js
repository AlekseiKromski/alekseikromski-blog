import {configureStore, getDefaultMiddleware} from '@reduxjs/toolkit'
import counterReducer from './store/shared'
import applicationReducer from "./store/application";

export default configureStore({
    reducer: {
        shared: counterReducer,
        application: applicationReducer
    },
    middleware: getDefaultMiddleware({
        serializableCheck: false,
    }),
})