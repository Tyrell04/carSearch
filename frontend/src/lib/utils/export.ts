import type { Car } from '../types';

export function exportCarsToCSV(cars: Car[], filename: string = 'saved-cars.csv'): void {
	if (cars.length === 0) {
		alert('Keine Autos zum Exportieren verfügbar');
		return;
	}

	// CSV headers
	const headers = ['HSN', 'TSN', 'Name', 'Hersteller', 'Erstellungsdatum'];
	
	// Convert cars to CSV rows
	const csvRows = cars.map(car => [
		car.HSN,
		car.TSN,
		`"${car.Name.replace(/"/g, '""')}"`, // Escape quotes in car names
		`"${car.Producer?.Name?.replace(/"/g, '""') || 'Unbekannt'}"`, // Escape quotes in producer names
		car.CreatedAt ? new Date(car.CreatedAt).toLocaleDateString('de-DE') : new Date().toLocaleDateString('de-DE')
	]);

	// Combine headers and data
	const csvContent = [headers, ...csvRows]
		.map(row => row.join(','))
		.join('\n');

	// Create and trigger download
	const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
	const link = document.createElement('a');
	
	if (link.download !== undefined) {
		const url = URL.createObjectURL(blob);
		link.setAttribute('href', url);
		link.setAttribute('download', filename);
		link.style.visibility = 'hidden';
		document.body.appendChild(link);
		link.click();
		document.body.removeChild(link);
		URL.revokeObjectURL(url);
	} else {
		// Fallback for older browsers
		window.open('data:text/csv;charset=utf-8,' + encodeURIComponent(csvContent));
	}
}
