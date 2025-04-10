import { Component, Input, Output, EventEmitter } from '@angular/core';
import { Doctor } from '../../../../core/models/doctor-interface';

@Component({
  selector: 'app-doctor-table',
  templateUrl: './doctor-table.component.html',
  styleUrl: './doctor-table.component.css'
})
export class DoctorTableComponent {
  @Input() doctors: Doctor[] = [];
  @Output() edit = new EventEmitter<number>();
  @Output() delete = new EventEmitter<number>();

  editDoctor(id: number) {
    this.edit.emit(id);
  }

  deleteDoctor(id: number) {
    this.delete.emit(id);
  }
}
