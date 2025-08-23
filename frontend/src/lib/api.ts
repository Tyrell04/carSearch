import type { Car, Producer, ApiResponse } from './types';

interface ApiConfig {
	baseUrl: string;
}

class ApiClient {
	private config: ApiConfig;

	constructor(config: ApiConfig) {
		this.config = config;
	}

	private async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
		const url = `${this.config.baseUrl}${endpoint}`;
		
		const response = await fetch(url, {
			headers: {
				'Content-Type': 'application/json',
				...options.headers,
			},
			...options,
		});

		const data = await response.json();

		if (!response.ok) {
			throw new Error(data.message || data.error || 'Request failed');
		}

		return data;
	}

	async searchCars(hsn: string, tsn?: string): Promise<ApiResponse> {
		const params = new URLSearchParams();
		params.set('hsn', hsn);
		if (tsn) params.set('tsn', tsn);

		return this.request<ApiResponse>(`/cars/search?${params.toString()}`);
	}

	// Future endpoints can be added here
	// async getCar(id: number): Promise<Car> { ... }
	// async saveCar(car: Car): Promise<Car> { ... }
}

// Environment-aware API client factory
function createApiClient(): ApiClient {
	const isDev = import.meta.env.DEV;
	const baseUrl = isDev 
		? 'http://localhost:3000/api' 
		: 'https://car.marc-schulz.online/api';

	return new ApiClient({ baseUrl });
}

export const api = createApiClient();
export type { ApiClient };
