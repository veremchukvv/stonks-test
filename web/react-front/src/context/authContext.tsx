import { createContext} from "react";

export const AuthContext = createContext({
    userName: '',
    isAuthenticated: false
})