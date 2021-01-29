import {RestClient} from "typed-rest-client/RestClient"
import {ChargeHistory, ChargeLog, Device} from "@/vals/vals";

interface GetDevicesResponse {
    devices: Array<Device>
}

interface GetChargeHistoryResponse {
    chargeHistory: ChargeHistory
}

// const rest = new RestClient("chargewatch-monitor", "http://localhost:8088")
const rest = new RestClient("chargewatch-monitor", "https://chargewatch.snowhork.com")

export async function GetDevices(userID: string): Promise<GetDevicesResponse> {
    const res = await rest.get<GetDevicesResponse>(`/user/${userID}/devices`)

    if(res.result == null) {
        throw new Error("Request Error")
    }

    return res.result
}

export async function GetChargeHistory(deviceID: string): Promise<GetChargeHistoryResponse> {
    const res = await rest.get<GetChargeHistoryResponse>(`/devices/${deviceID}/chargeHistory`)

    if(res.result == null) {
        throw new Error("Request Error")
    }

    return res.result
}