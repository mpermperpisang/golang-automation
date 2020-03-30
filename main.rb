# frozen_string_literal: true

require "report_builder"

options = {
  report_title: 'Smoke Testing Report',
  report_types: ['json', 'html']
}

ReportBuilder.build_report options



