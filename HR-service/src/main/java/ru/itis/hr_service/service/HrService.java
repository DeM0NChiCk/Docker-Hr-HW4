package ru.itis.hr_service.service;

import ru.itis.hr_service.dto.JobCvPair;

import java.util.List;

public interface HrService {
    List<JobCvPair> getPairs();
}
