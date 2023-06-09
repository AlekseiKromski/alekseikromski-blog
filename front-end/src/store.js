import { configureStore } from '@reduxjs/toolkit'
import counterReducer from './store/shared'

export default configureStore({
    reducer: {
        shared: counterReducer
    }
})