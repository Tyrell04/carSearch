export interface Producer {
	ID?: number;
	HSN: string;
	Name: string;
	CreatedAt?: string;
	UpdatedAt?: string;
}

export interface Car {
	ID?: number;
	HSN: string;
	TSN: string;
	Name: string;
	Producer?: Producer;
	CreatedAt?: string;
	UpdatedAt?: string;
}

export interface ApiResponse {
	message?: string;
	error?: string;
	cars?: Car[];
	car?: Car;
	producer?: Producer;
}
