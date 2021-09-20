import { MCLService } from "./rpc.gen.js"

const RPCService = new MCLService('', window.fetch.bind(window))

export default RPCService
