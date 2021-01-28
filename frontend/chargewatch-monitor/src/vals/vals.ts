export interface Device {
  id: string
  userID: string
  name: string
  description: string
  state: string
  charge: Charge
}

export interface Charge {
  value: number
  charging: boolean
  timestamp: number
}

export interface ChargeHistory {
  deviceID: string
  logs: Array<ChargeLog>
}


export interface ChargeLog {
  deviceID: string
  charge: Charge
}
