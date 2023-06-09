import { createSlice } from '@reduxjs/toolkit'

export const sharedSlice = createSlice({
    name: 'sharedSlice',
    initialState: {
        categories: []
    },
    reducers: {
        importCategories: (state,categories) => {
            state.categories = [...categories.payload]
        },
    }
})

// Action creators are generated for each case reducer function
export const { importCategories } = sharedSlice.actions

export default sharedSlice.reducer