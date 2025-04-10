import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
// import { AdminDashboardComponent } from './pages/admin-dashboard/admin-dashboard.component';
import { AdminDashboardComponent } from './pages/admin-dashboard/admin-dashboard/admin-dashboard.component';
import { AdminRoutingModule } from './admin-routing.module';
// import { AdminRoutingModule } from './admin-routing.module';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [AdminDashboardComponent],
  imports: [
    CommonModule,
    AdminRoutingModule,
    FormsModule
  ],
  exports:[
    AdminDashboardComponent
  ]
})
export class AdminModule { }
