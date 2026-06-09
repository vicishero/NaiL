declare namespace Api {

    namespace Auth {

        interface Api {
            post: {
                /** 用户登录 */
                login: (params: NetParams.AuthUserLogin) => Promise<NetReq.AuthUserLogin>;
                /** 用户注册 */
                register: (params: NetParams.AuthUserRegister) => Promise<NetReq.AuthUserRegister>;
                /** 获取钱包登录nonce */
                walletNonce: (params: NetParams.AuthWalletNonce) => Promise<NetReq.AuthWalletNonce>;
                /** 钱包登录 */
                walletLogin: (params: NetParams.AuthWalletLogin) => Promise<NetReq.AuthWalletLogin>;
            };
        }

        namespace NetParams {
            interface AuthUserLogin {
                /** 用户名 */
                username: string;
                /** 密码 */
                password: string;
            }

            interface AuthUserRegister extends AuthUserLogin {

            }

            interface AuthWalletNonce {
                /** 钱包地址 */
                address: string;
            }

            interface AuthWalletLogin {
                /** 钱包地址 */
                address: string;
                /** 签名 */
                signature: string;
                /** nonce */
                nonce: string;
            }
        }

        namespace NetReq {
            interface AuthUserLogin {
                token: string;
            }

            interface AuthUserRegister {
                /** 用户UID */
                id: string;
                /** 用户名 */
                username: string;
            }

            interface AuthWalletNonce {
                nonce: string;
                message: string;
            }

            interface AuthWalletLogin {
                token: string;
                is_new_user: boolean;
            }
        }

    }

}